package fronted

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCaching(t *testing.T) {
	dir := t.TempDir()
	cacheFile := filepath.Join(dir, "cachefile.1")

	cloudsackID := "cloudsack"

	providers := map[string]*Provider{
		testProviderID: NewProvider(nil, "", nil, nil, nil, nil, ""),
		cloudsackID:    NewProvider(nil, "", nil, nil, nil, nil, ""),
	}

	log.Debug("Creating fronted")
	makeFronted := func() *fronted {
		f := &fronted{
			fronts:              newThreadSafeFronts(1000),
			maxAllowedCachedAge: 250 * time.Millisecond,
			maxCacheSize:        4,
			cacheSaveInterval:   50 * time.Millisecond,
			cacheDirty:          make(chan interface{}, 1),
			cacheClosed:         make(chan interface{}),
			providers:           providers,
			stopCh:              make(chan interface{}, 10),
			defaultProviderID:   cloudsackID,
		}
		go f.maintainCache(cacheFile)
		return f
	}

	now := time.Now()
	mb := &front{Masquerade: Masquerade{Domain: "b", IpAddress: "2"}, LastSucceeded: now, ProviderID: testProviderID}
	mc := &front{Masquerade: Masquerade{Domain: "c", IpAddress: "3"}, LastSucceeded: now, ProviderID: ""}         // defaulted
	md := &front{Masquerade: Masquerade{Domain: "d", IpAddress: "4"}, LastSucceeded: now, ProviderID: "sadcloud"} // skipped

	f := makeFronted()

	log.Debug("Adding fronts")
	f.fronts.fronts = append(f.fronts.fronts, mb, mc, md)

	readCached := func() []*front {
		log.Debug("Reading cached fronts")
		var result []*front
		b, err := os.ReadFile(cacheFile)
		require.NoError(t, err, "Unable to read cache file")
		err = json.Unmarshal(b, &result)
		require.NoError(t, err, "Unable to unmarshal cache file")
		return result
	}

	// Save the cache
	f.markCacheDirty()

	time.Sleep(f.cacheSaveInterval * 2)
	f.Close()

	time.Sleep(50 * time.Millisecond)

	log.Debug("Reopening fronted")
	// Reopen cache file and make sure right data was in there
	f = makeFronted()
	f.prepopulateFronts(cacheFile)
	masquerades := readCached()
	require.Len(t, masquerades, 3, "Wrong number of masquerades read")
	for i, expected := range []*front{mb, mc, md} {
		require.Equal(t, expected.Domain, masquerades[i].Domain, "Wrong masquerade at position %d", i)
		require.Equal(t, expected.IpAddress, masquerades[i].IpAddress, "Masquerade at position %d has wrong IpAddress", 0)
		require.Equal(t, expected.ProviderID, masquerades[i].ProviderID, "Masquerade at position %d has wrong ProviderID", 0)
		require.Equal(t, now.Unix(), masquerades[i].LastSucceeded.Unix(), "Masquerade at position %d has wrong LastSucceeded", 0)
	}
	f.Close()
}
