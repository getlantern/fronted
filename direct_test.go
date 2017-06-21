package fronted

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testURL = "http://d157vud77ygy87.cloudfront.net/measurements"
)

func TestDirectDomainFronting(t *testing.T) {
	dir, err := ioutil.TempDir("", "direct_test")
	if !assert.NoError(t, err, "Unable to create temp dir") {
		return
	}
	defer os.RemoveAll(dir)
	cacheFile := filepath.Join(dir, "cachefile")
	doTestDomainFronting(t, cacheFile)
	time.Sleep(defaultCacheSaveInterval * 2)
	// Then try again, this time reusing the existing cacheFile
	doTestDomainFronting(t, cacheFile)
}

func doTestDomainFronting(t *testing.T, cacheFile string) {
	ConfigureCachingForTest(t, cacheFile)
	client := &http.Client{
		Transport: NewDirect(30 * time.Second),
	}
	req, _ := http.NewRequest(http.MethodPost, testURL, strings.NewReader("{'bad': 'stuff'}"))
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not get response: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Unexpected response status: %v", resp.StatusCode)
	}

	log.Debugf("DIRECT DOMAIN FRONTING TEST SUCCEEDED")
}
