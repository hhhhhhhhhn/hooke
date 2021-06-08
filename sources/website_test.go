package sources

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dummyDownloadWebsite(url string) string {
	return url
}

func TestDownloadWebsite(t *testing.T) {
	// FLAKY
	downloadWebsite = actualDownloadWebsite
	response := downloadWebsite("https://wikipedia.org")
	assert.True(t, strings.Contains(response, "Wikipedia"))
}

func TestParseHtml(t *testing.T) {
	testHtml := "<script>should be deleted</script><p>should not</p>"
	assert.Equal(t, "should not", parseHtml(testHtml))
}

func TestWebsite(t *testing.T) {
	downloadWebsite = dummyDownloadWebsite

	testText := Website("<script>nothing</script><p>test url</p>", &testLanguage)

	assert.Equal(t, "test url", testText.Raw)
}
