package fronted

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type cacheOp struct {
	m      masquerade
	remove bool
	close  bool
}

func (d *direct) initCaching(cacheFile string) int {
	cache := d.prepopulateMasquerades(cacheFile)
	prevetted := len(cache)
	go d.fillCache(cache, cacheFile)
	return prevetted
}

func (d *direct) prepopulateMasquerades(cacheFile string) []masquerade {
	var cache []masquerade
	bytes, err := ioutil.ReadFile(cacheFile)
	if err == nil {
		log.Debugf("Attempting to prepopulate masquerades from cache")
		var masquerades []masquerade
		err := json.Unmarshal(bytes, &masquerades)
		if err != nil {
			log.Errorf("Error prepopulating cached masquerades: %v", err)
			return cache
		}

		log.Debugf("Cache contained %d masquerades", len(masquerades))
		now := time.Now()
		for _, m := range masquerades {
			if now.Sub(m.LastVetted) < d.maxAllowedCachedAge {
				// fill in default for masquerades lacking provider id
				if m.ProviderID == "" {
					m.ProviderID = d.defaultProviderID
				}
				// Skip entries for providers that are not configured.
				_, ok := d.providers[m.ProviderID]
				if !ok {
					log.Debugf("Skipping cached entry for unknown/disabled provider %s", m.ProviderID)
					continue
				}
				select {
				case d.cached <- m:
					// submitted
					cache = append(cache, m)
				default:
					// channel full, that's okay
				}
			}
		}
	}

	return cache
}

func (d *direct) fillCache(cache []masquerade, cacheFile string) {
	saveTicker := time.NewTicker(d.cacheSaveInterval)
	defer saveTicker.Stop()
	cacheChanged := false
	for {
		select {
		case op := <-d.toCache:
			if op.close {
				log.Debug("Cache closed, stop filling")
				return
			}
			m := op.m
			if op.remove {
				newCache := make([]masquerade, len(cache))
				for _, existing := range cache {
					if existing.Domain == m.Domain && existing.IpAddress == m.IpAddress {
						log.Debugf("Removing masquerade for %v (%v)", m.Domain, m.IpAddress)
					} else {
						newCache = append(newCache, existing)
					}
				}
				cache = newCache
			} else {
				log.Debugf("Caching vetted masquerade for %v (%v)", m.Domain, m.IpAddress)
				cache = append(cache, m)
			}
			cacheChanged = true
		case <-saveTicker.C:
			if !cacheChanged {
				continue
			}
			log.Debug("Saving updated masquerade cache")
			// Truncate cache to max length if necessary
			if len(cache) > d.maxCacheSize {
				truncated := make([]masquerade, d.maxCacheSize)
				copy(truncated, cache[len(cache)-d.maxCacheSize:])
				cache = truncated
			}
			b, err := json.Marshal(cache)
			if err != nil {
				log.Errorf("Unable to marshal cache to JSON: %v", err)
				break
			}
			err = ioutil.WriteFile(cacheFile, b, 0644)
			if err != nil {
				log.Errorf("Unable to save cache to disk: %v", err)
			}
			cacheChanged = false
		}
	}
}

func (d *direct) closeCache() {
	d.toCache <- &cacheOp{close: true}
}
