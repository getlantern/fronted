package fronted

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type masqueradeCache struct {
	filename       string
	maxSize        int
	maxAge         time.Duration
	newEntries     []masquerade
	newEntriesLock sync.Mutex
	done           chan struct{}
	closeOnce      sync.Once
}

func newMasqueradeCache(filename string, maxSize int, maxAge, saveInterval time.Duration) *masqueradeCache {
	c := &masqueradeCache{
		filename, maxSize, maxAge, []masquerade{}, sync.Mutex{}, make(chan struct{}), sync.Once{},
	}
	go func() {
		ticker := time.NewTicker(saveInterval)
		defer ticker.Stop()
		for {
			select {
			case <-c.done:
				// Flush to disk.
				if err := c.saveNewEntries(); err != nil {
					log.Errorf("save routine encountered error saving while closing: %v", err)
				}
				return
			case <-ticker.C:
				if err := c.saveNewEntries(); err != nil {
					log.Errorf("save routine encountered error: %v", err)
				}
			}
		}
	}()
	return c
}

func (c *masqueradeCache) read() ([]masquerade, error) {
	c.newEntriesLock.Lock()
	defer c.newEntriesLock.Unlock()
	return c.readUnsafe()

}

func (c *masqueradeCache) readUnsafe() ([]masquerade, error) {
	m := make([]masquerade, len(c.newEntries))
	copy(m, c.newEntries)

	if _, err := os.Stat(c.filename); errors.Is(err, os.ErrNotExist) {
		return m, nil
	}
	f, err := os.Open(c.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache file (%s) for reading: %w", c.filename, err)
	}
	defer f.Close()
	fromDisk := []masquerade{}
	if err := json.NewDecoder(f).Decode(&fromDisk); err != nil {
		return nil, fmt.Errorf("failed to decode cache file: %w", err)
	}
	for _, masq := range fromDisk {
		if time.Since(masq.LastVetted) < c.maxAge {
			m = append(m, masq)
		}
	}
	return m, nil
}

func (c *masqueradeCache) write(m masquerade) {
	select {
	case <-c.done:
		// No-op if the cache is closed.
	default:
		c.newEntriesLock.Lock()
		c.newEntries = append(c.newEntries, m)
		c.newEntriesLock.Unlock()
	}
}

func (c *masqueradeCache) saveNewEntries() error {
	c.newEntriesLock.Lock()
	defer c.newEntriesLock.Unlock()
	if len(c.newEntries) == 0 {
		return nil
	}
	current, err := c.readUnsafe()
	if err != nil {
		return fmt.Errorf("failed to read current entries: %w", err)
	}
	if len(current) > c.maxSize {
		current = current[:c.maxSize]
	}
	b, err := json.Marshal(current)
	if err != nil {
		return fmt.Errorf("failed to marshal entries as JSON: %w", err)
	}
	if err := ioutil.WriteFile(c.filename, b, 0644); err != nil {
		return fmt.Errorf("failed to write updates to disk: %w", err)
	}
	c.newEntries = []masquerade{}
	return nil
}

func (c *masqueradeCache) close() {
	log.Debugf("cache at %s closed", c.filename)
	c.closeOnce.Do(func() { close(c.done) })
}

// We track open caches centrally to:
//	(a) Ensure only a single masqueradeCache is open for each file at any time. Otherwise, different
// 	    instances may clobber each other.
//	(b) Ensure this single masqueradeCache instance is closed when all users are done using it.

var globalCacheManager = newCacheManager()

type cacheManager struct {
	// filename -> cache
	openCaches map[string]*masqueradeCache

	// filename -> open handles
	openHandles map[string]int

	sync.Mutex
}

func newCacheManager() *cacheManager {
	return &cacheManager{map[string]*masqueradeCache{}, map[string]int{}, sync.Mutex{}}
}

func (cm *cacheManager) get(filename string, maxSize int, maxAge, saveInterval time.Duration) *masqueradeCache {
	cm.Lock()
	defer cm.Unlock()
	c, ok := cm.openCaches[filename]
	if !ok {
		c = newMasqueradeCache(filename, maxSize, maxAge, saveInterval)
		cm.openCaches[filename] = c
		cm.openHandles[filename] = 0
	}
	cm.openHandles[filename]++
	return c
}

func (cm *cacheManager) closeHandle(c *masqueradeCache) {
	cm.Lock()
	defer cm.Unlock()
	c, ok := cm.openCaches[c.filename]
	if !ok {
		return
	}
	cm.openHandles[c.filename]--
	if cm.openHandles[c.filename] <= 0 {
		c.close()
		delete(cm.openCaches, c.filename)
		delete(cm.openHandles, c.filename)
	}
}
