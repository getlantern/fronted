package fronted

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCaching(t *testing.T) {
	dir, err := os.MkdirTemp("", "direct_test")
	if !assert.NoError(t, err, "Unable to create temp dir") {
		return
	}
	defer os.RemoveAll(dir)
	cacheFile := filepath.Join(dir, "cachefile.1")

	cloudsackID := "cloudsack"

	providers := map[string]*Provider{
		testProviderID: NewProvider(nil, "", nil, nil, nil, nil, nil),
		cloudsackID:    NewProvider(nil, "", nil, nil, nil, nil, nil),
	}

	makeFronted := func() *fronted {
		f := &fronted{
			fronts:              make(sortedFronts, 0, 1000),
			maxAllowedCachedAge: 250 * time.Millisecond,
			maxCacheSize:        4,
			cacheSaveInterval:   50 * time.Millisecond,
			cacheDirty:          make(chan interface{}, 1),
			cacheClosed:         make(chan interface{}),
			providers:           providers,
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
	f.fronts = append(f.fronts, mb, mc, md)

	readCached := func() []*front {
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

	// Reopen cache file and make sure right data was in there
	f = makeFronted()
	f.prepopulateMasquerades(cacheFile)
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
