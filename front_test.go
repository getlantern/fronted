package fronted

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProvider(t *testing.T) {
	verifyHostname := "verifyHostname.com"
	var tests = []struct {
		name                string
		givenHosts          map[string]string
		givenTestURL        string
		givenMasquerades    []*Masquerade
		givenValidator      ResponseValidator
		givenPassthrough    []string
		givenSNIConfig      *SNIConfig
		givenVerifyHostname *string
		assert              func(t *testing.T, actual *Provider)
	}{
		{
			name:         "should return a new provider without host aliases, masquerades and passthrough",
			givenHosts:   map[string]string{},
			givenTestURL: "http://test.com",
			assert: func(t *testing.T, actual *Provider) {
				assert.Empty(t, actual.HostAliases)
				assert.Empty(t, actual.Masquerades)
				assert.Empty(t, actual.PassthroughPatterns)
				assert.Equal(t, "http://test.com", actual.TestURL)
				assert.Nil(t, actual.Validator)
				assert.Nil(t, actual.SNIConfig)
			},
		},
		{
			name:             "should return a new provider with host aliases, masquerades and passthrough",
			givenHosts:       map[string]string{"host1": "alias1", "host2": "alias2"},
			givenTestURL:     "http://test.com",
			givenMasquerades: []*Masquerade{{Domain: "domain1", IpAddress: "127.0.0.1"}},
			givenValidator:   func(*http.Response) error { return nil },
			givenPassthrough: []string{"passthrough1", "passthrough2"},
			givenSNIConfig: &SNIConfig{
				UseArbitrarySNIs: true,
				ArbitrarySNIs:    []string{"sni1.com", "sni2.com"},
			},
			givenVerifyHostname: &verifyHostname,
			assert: func(t *testing.T, actual *Provider) {
				assert.Equal(t, "http://test.com", actual.TestURL)
				assert.Equal(t, "alias1", actual.HostAliases["host1"])
				assert.Equal(t, "alias2", actual.HostAliases["host2"])
				assert.Equal(t, 1, len(actual.Masquerades))
				assert.Equal(t, "domain1", actual.Masquerades[0].Domain)
				assert.Equal(t, "127.0.0.1", actual.Masquerades[0].IpAddress)
				assert.Equal(t, "sni1.com", actual.Masquerades[0].SNI)
				assert.Equal(t, verifyHostname, *actual.Masquerades[0].VerifyHostname)
				assert.Equal(t, 2, len(actual.PassthroughPatterns))
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := NewProvider(tt.givenHosts, tt.givenTestURL, tt.givenMasquerades, tt.givenValidator, tt.givenPassthrough, tt.givenSNIConfig, tt.givenVerifyHostname)
			tt.assert(t, actual)
		})
	}
}

func TestGenerateSNI(t *testing.T) {
	emptyMasquerade := new(Masquerade)
	var tests = []struct {
		name            string
		assert          func(t *testing.T, actual string)
		givenConfig     *SNIConfig
		givenMasquerade *Masquerade
	}{
		{
			name:            "should return a empty string when given SNI config is nil",
			givenConfig:     nil,
			givenMasquerade: emptyMasquerade,
			assert: func(t *testing.T, actual string) {
				assert.Empty(t, actual)
			},
		},
		{
			name: "should return a empty string when given SNI config is not nil and UseArbitrarySNIs is false",
			givenConfig: &SNIConfig{
				UseArbitrarySNIs: false,
			},
			givenMasquerade: emptyMasquerade,
			assert: func(t *testing.T, actual string) {
				assert.Empty(t, actual)
			},
		},
		{
			name: "should return a empty SNI when the list of arbitrary SNIs is empty",
			givenConfig: &SNIConfig{
				UseArbitrarySNIs: true,
				ArbitrarySNIs:    []string{},
			},
			givenMasquerade: &Masquerade{
				IpAddress: "1.1.1.1",
				Domain:    "randomdomain.net",
			},
			assert: func(t *testing.T, actual string) {
				assert.Empty(t, actual)
			},
		},
		{
			name: "should return a SNI when given SNI config is not nil and UseArbitrarySNIs is true",
			givenConfig: &SNIConfig{
				UseArbitrarySNIs: true,
				ArbitrarySNIs:    []string{"sni1.com", "sni2.com"},
			},
			givenMasquerade: &Masquerade{
				IpAddress: "1.1.1.1",
				Domain:    "randomdomain.net",
			},
			assert: func(t *testing.T, actual string) {
				assert.NotEmpty(t, actual)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := generateSNI(tt.givenConfig, tt.givenMasquerade)
			tt.assert(t, actual)
		})
	}
}
