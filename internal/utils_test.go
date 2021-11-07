package internal

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const expectedURL = "https://mapped.chargify.com/customers/lookup.json?reference=ORGOberbrunnerStanton"

func TestUrlEncoding(t *testing.T) {
	u, err := url.Parse(expectedURL)
	assert.NoError(t, err)
	assert.NotNil(t, u)

	url := u.String()
	assert.Equal(t, expectedURL, url)

}

func TestJoinUrl(t *testing.T) {
	baseUrl := "https://mapped.chargify.com"
	path := "customers/lookup.json?reference=ORGOberbrunnerStanton"

	url, err := ResolveEndpointUrl(baseUrl, path)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, url.String())
}
