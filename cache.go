package fronted

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func (d *fronted) initCaching(cacheFile string) {
	d.prepopulateFronts(cacheFile)
	go d.maintainCache(cacheFile)
}

func (d *fronted) prepopulateFronts(cacheFile string) {
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
	var cachedFronts []*front
	if err := json.Unmarshal(bytes, &cachedFronts); err != nil {
		log.Errorf("Error reading cached masquerades: %v", err)
		return
	}

	log.Debugf("Cache contained %d masquerades", len(cachedFronts))
	now := time.Now()

	// update last succeeded status of masquerades based on cached values
	for _, f := range d.fronts {
		for _, cf := range cachedFronts {
			sameFront := cf.ProviderID == f.getProviderID() && cf.Domain == f.getDomain() && cf.IpAddress == f.getIpAddress()
			cachedValueFresh := now.Sub(f.lastSucceeded()) < d.maxAllowedCachedAge
			if sameFront && cachedValueFresh {
				f.setLastSucceeded(cf.LastSucceeded)
			}
		}
	}
}

func (d *fronted) markCacheDirty() {
	select {
	case d.cacheDirty <- nil:
		// okay
	default:
		// already dirty
	}
}

func (d *fronted) maintainCache(cacheFile string) {
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

func (d *fronted) updateCache(cacheFile string) {
	log.Debugf("Updating cache at %v", cacheFile)
	cache := d.fronts.sortedCopy()
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
		// Log the directory of the cache file and if it exists for debugging purposes
		parent := filepath.Dir(cacheFile)
		// check if the parent directory exists
		if _, err := os.Stat(parent); err == nil {
			// parent directory exists
			log.Debugf("Parent directory of cache file exists: %v", parent)
		} else {
			// parent directory does not exist
			log.Debugf("Parent directory of cache file does not exist: %v", parent)
		}
	} else {
		log.Debugf("Cache saved to disk")
	}
}
