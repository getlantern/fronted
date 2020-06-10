package fronted

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCaching(t *testing.T) {
	const (
		maxSize      = 3
		maxAge       = 250 * time.Millisecond
		saveInterval = 50 * time.Millisecond
		timeout      = time.Second
	)

	dir, err := ioutil.TempDir("", "direct_test")
	require.NoError(t, err, "Unable to create temp dir")
	defer os.RemoveAll(dir)
	cacheFile := filepath.Join(dir, "cachefile.1")

	cloudsackID := "cloudsack"

	providers := map[string]*Provider{
		testProviderID: NewProvider(nil, "", nil, nil, nil),
		cloudsackID:    NewProvider(nil, "", nil, nil, nil),
	}

	cache, err := newMasqueradeCache(cacheFile, maxSize, maxAge, saveInterval)
	require.NoError(t, err)

	makeDirect := func() *direct {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		d, err := newDirect(ctx, providers, cloudsackID, cache, DirectOptions{})
		require.NoError(t, err)
		d.candidates = make(chan masquerade, 1000)
		d.masquerades = make(chan masquerade, 1000)
		return d
	}

	now := time.Now()
	cache.write(masquerade{Masquerade{Domain: "a", IpAddress: "1"}, now, testProviderID})
	cache.write(masquerade{Masquerade{Domain: "b", IpAddress: "2"}, now, testProviderID})
	cache.write(masquerade{Masquerade{Domain: "c", IpAddress: "3"}, now, ""})         // defaulted
	cache.write(masquerade{Masquerade{Domain: "d", IpAddress: "4"}, now, "sadcloud"}) // skipped

	d := makeDirect()

	readMasquerades := func() []masquerade {
		var result []masquerade
		for {
			select {
			case m := <-d.masquerades:
				result = append(result, m)
			default:
				return result
			}
		}
	}

	// Fill the cache
	time.Sleep(saveInterval * 2)
	cache.close()
	time.Sleep(saveInterval)

	// Reopen cache file and make sure right data was in there
	cache, err = newMasqueradeCache(cacheFile, maxSize, maxAge, saveInterval)
	require.NoError(t, err)
	d = makeDirect()
	masquerades := readMasquerades()
	assert.Len(t, masquerades, 2, "Wrong number of masquerades read")
	assert.Equal(t, "b", masquerades[0].Domain, "Wrong masquerade at position 0")
	assert.Equal(t, "2", masquerades[0].IpAddress, "Masquerade at position 0 has wrong IpAddress")
	assert.Equal(t, testProviderID, masquerades[0].ProviderID, "Masquerade at position 0 has wrong ProviderID")
	assert.Equal(t, "c", masquerades[1].Domain, "Wrong masquerade at position 0")
	assert.Equal(t, "3", masquerades[1].IpAddress, "Masquerade at position 1 has wrong IpAddress")
	assert.Equal(t, cloudsackID, masquerades[1].ProviderID, "Masquerade at position 1 has wrong ProviderID")
	cache.close()

	time.Sleep(maxAge)
	cache, err = newMasqueradeCache(cacheFile, maxSize, maxAge, saveInterval)
	require.NoError(t, err)
	d = makeDirect()
	assert.Empty(t, readMasquerades(), "Cache should be empty after masquerades expire")
	cache.close()
}
