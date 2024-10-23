package fronted

import (
	"encoding/json"
	"os"
	"time"
)

func (d *direct) initCaching(cacheFile string) {
	d.prepopulateMasquerades(cacheFile)
	go d.maintainCache(cacheFile)
}

func (d *direct) prepopulateMasquerades(cacheFile string) {
	bytes, err := os.ReadFile(cacheFile)
	if err != nil {
		// This is not a big deal since we'll just fill the cache later
		log.Debugf("ignorable error: Unable to read cache file for prepopulation: %v", err)
		return
	}

	if len(bytes) == 0 {
		// This can happen if the file is empty or just not there
		log.Debug("ignorable error: Cache file is empty")
		return
	}

	log.Debugf("Attempting to prepopulate masquerades from cache file: %v", cacheFile)
	var cachedMasquerades []*masquerade
	if err := json.Unmarshal(bytes, &cachedMasquerades); err != nil {
		log.Errorf("Error reading cached masquerades: %v", err)
		return
	}

	log.Debugf("Cache contained %d masquerades", len(cachedMasquerades))
	now := time.Now()

	// update last succeeded status of masquerades based on cached values
	for _, m := range d.masquerades {
		for _, cm := range cachedMasquerades {
			sameMasquerade := cm.ProviderID == m.getProviderID() && cm.Domain == m.getDomain() && cm.IpAddress == m.getIpAddress()
			cachedValueFresh := now.Sub(m.lastSucceeded()) < d.maxAllowedCachedAge
			if sameMasquerade && cachedValueFresh {
				m.setLastSucceeded(cm.LastSucceeded)
			}
		}
	}
}

func (d *direct) markCacheDirty() {
	select {
	case d.cacheDirty <- nil:
		// okay
	default:
		// already dirty
	}
}

func (d *direct) maintainCache(cacheFile string) {
	for {
		select {
		case <-d.cacheClosed:
			return
		case <-time.After(d.cacheSaveInterval):
			select {
			case <-d.cacheClosed:
				return
			case <-d.cacheDirty:
				d.updateCache(cacheFile)
			}
		}
	}
}

func (d *direct) updateCache(cacheFile string) {
	log.Debugf("Updating cache at %v", cacheFile)
	cache := d.masquerades.sortedCopy()
	sizeToSave := len(cache)
	if d.maxCacheSize < sizeToSave {
		sizeToSave = d.maxCacheSize
	}
	b, err := json.Marshal(cache[:sizeToSave])
	if err != nil {
		log.Errorf("Unable to marshal cache to JSON: %v", err)
		return
	}
	err = os.WriteFile(cacheFile, b, 0644)
	if err != nil {
		log.Errorf("Unable to save cache to disk: %v", err)
	}
}

func (d *direct) closeCache() {
	d.closeCacheOnce.Do(func() {
		close(d.cacheClosed)
	})
}
