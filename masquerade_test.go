package fronted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
