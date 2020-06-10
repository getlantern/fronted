package fronted

import (
	"encoding/json"
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

func newMasqueradeCache(
	filename string, maxSize int, maxAge, saveInterval time.Duration) (*masqueradeCache, error) {

	_, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to stat cache file: %w", err)
	}
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
	return c, nil
}

func (c *masqueradeCache) read() ([]masquerade, error) {
	f, err := os.Open(c.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache file (%s) for reading: %w", c.filename, err)
	}
	defer f.Close()
	_m := []masquerade{}
	if err := json.NewDecoder(f).Decode(&_m); err != nil {
		return nil, fmt.Errorf("failed to decode cache file: %w", err)
	}
	m := []masquerade{}
	for _, masq := range _m {
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
	current, err := c.read()
	if err != nil {
		return fmt.Errorf("failed to read current entries: %w", err)
	}
	current = append(current, c.newEntries...)
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
	return nil
}

func (c *masqueradeCache) close() {
	log.Debugf("cache at %s closed", c.filename)
	c.closeOnce.Do(func() { close(c.done) })
}
