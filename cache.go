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
			log.Error("Error creating cache directory", "error", err)
			return
		}
	}
	if err != nil {
		log.Error("Error reading cache file", "error", err)
		return
	}

	if len(bytes) == 0 {
		// This can happen if the file is empty or just not there
		log.Debug("ignorable error: Cache file is empty")
		return
	}

	log.Debug("Attempting to prepopulate masquerades from cache file", "cacheFile", cacheFile)
	var cachedFronts []*front
	if err := json.Unmarshal(bytes, &cachedFronts); err != nil {
		log.Error("Error reading cached masquerades", "error", err)
		return
	}

	log.Debug("Found masquerades in cache", "count", len(cachedFronts))
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
	log.Debug("Updating cache", "cacheFile", cacheFile)
	cache := f.fronts.sortedCopy()
	sizeToSave := len(cache)
	if f.maxCacheSize < sizeToSave {
		sizeToSave = f.maxCacheSize
	}
	b, err := json.Marshal(cache[:sizeToSave])
	if err != nil {
		log.Error("Unable to marshal cache to JSON", "error", err)
		return
	}
	err = os.WriteFile(cacheFile, b, 0644)
	if err != nil {
		log.Error("Unable to save cache to disk", "error", err)
		// Log the directory of the cache file and if it exists for debugging purposes
		parent := filepath.Dir(cacheFile)
		// check if the parent directory exists
		if _, err := os.Stat(parent); err == nil {
			// parent directory exists
			log.Debug("Parent directory of cache file exists", "directory", parent)
		} else {
			// parent directory does not exist
			log.Debug("Parent directory of cache file does not exist", "directory", parent)
		}
	} else {
		log.Debug("Cache saved to disk")
	}
}
