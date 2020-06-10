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

	cache := newMasqueradeCache(cacheFile, maxSize, maxAge, saveInterval)
	makeDirect := func() *direct {
		d, err := newDirect(context.Background(), providers, cloudsackID, 0, cache, DirectOptions{})
		require.NoError(t, err)
		d.candidates = make(chan masquerade, 1000)
		d.masquerades = make(chan masquerade, 1000)
		d.initFromCache()
		return d
	}

	// Fill the cache
	now := time.Now()
	cache.write(masquerade{Masquerade{Domain: "a", IpAddress: "1"}, now, testProviderID})
	cache.write(masquerade{Masquerade{Domain: "b", IpAddress: "2"}, now, testProviderID})
	cache.write(masquerade{Masquerade{Domain: "c", IpAddress: "3"}, now, ""})         // defaulted
	cache.write(masquerade{Masquerade{Domain: "d", IpAddress: "4"}, now, "sadcloud"}) // skipped
	time.Sleep(saveInterval * 2)

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

	cache.close()
	time.Sleep(saveInterval)

	// Reopen cache file and make sure right data was in there
	cache = newMasqueradeCache(cacheFile, maxSize, maxAge, saveInterval)
	d = makeDirect()
	masquerades := readMasquerades()
	assert.Len(t, masquerades, 3, "Wrong number of masquerades read")
	for i, expected := range []struct {
		domain, ip, providerID string
	}{
		{"a", "1", testProviderID},
		{"b", "2", testProviderID},
		{"c", "3", cloudsackID},
	} {
		t.Logf("checking masquerade %d", i)
		assert.Equal(t, expected.domain, masquerades[i].Domain)
		assert.Equal(t, expected.ip, masquerades[i].IpAddress)
		assert.Equal(t, expected.providerID, masquerades[i].ProviderID)
	}
	cache.close()

	time.Sleep(maxAge)
	cache = newMasqueradeCache(cacheFile, maxSize, maxAge, saveInterval)
	require.NoError(t, err)
	d = makeDirect()
	assert.Empty(t, readMasquerades(), "Cache should be empty after masquerades expire")
	cache.close()
}
