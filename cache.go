package fronted

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func (f *fronted) initCaching(cacheFile string) {
	f.prepopulateFronts(cacheFile)
	go f.maintainCache(cacheFile)
}

func (f *fronted) prepopulateFronts(cacheFile string) {
	bytes, err := os.ReadFile(cacheFile)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(filepath.Dir(cacheFile), 0755); err != nil {
			log.Errorf("Error creating cache directory: %v", err)
			return
		}
	}
	if err != nil {
		log.Errorf("Error reading cache file: %v", err)
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
	for _, fr := range f.fronts.fronts {
		for _, cf := range cachedFronts {
			sameFront := cf.ProviderID == fr.getProviderID() && cf.Domain == fr.getDomain() && cf.IpAddress == fr.getIpAddress()
			cachedValueFresh := now.Sub(fr.lastSucceeded()) < f.maxAllowedCachedAge
			if sameFront && cachedValueFresh {
				fr.setLastSucceeded(cf.LastSucceeded)
			}
		}
	}
}

func (f *fronted) markCacheDirty() {
	select {
	case f.cacheDirty <- nil:
		// okay
	default:
		// already dirty
	}
}

func (f *fronted) maintainCache(cacheFile string) {
	for {
		select {
		case <-f.cacheClosed:
			return
		case <-time.After(f.cacheSaveInterval):
			select {
			case <-f.cacheClosed:
				return
			case <-f.cacheDirty:
				f.updateCache(cacheFile)
			}
		}
	}
}

func (f *fronted) updateCache(cacheFile string) {
	log.Debugf("Updating cache at %v", cacheFile)
	cache := f.fronts.sortedCopy()
	sizeToSave := min(f.maxCacheSize, len(cache))
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
