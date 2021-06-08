package sources

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hhhhhhhhhn/hooke"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, min(1, 2))
	assert.Equal(t, 1, min(2, 1))
}

func TestDivideIntoParts(t *testing.T) {
	testArray := []string{"0", "1", "2", "3", "4"}
	expected := [][]string{
		{"0", "1"},
		{"2", "3"},
		{"4"},
	}
	assert.Equal(t, expected, divideIntoParts(testArray, 2))
}

func TestElementInArray(t *testing.T) {
	testArray := []string{"0", "1", "2", "3"}
	assert.Equal(t, false, elementInArray(testArray, "a"))
	assert.Equal(t, true, elementInArray(testArray, "1"))
}

func TestAppendNonDuplicates(t *testing.T) {
	testArray := []string{"0", "2", "4", "8", "10"}
	testAppended := []string{"1", "2", "3", "4"}
	expected := []string{"0", "2", "4", "8", "10", "1", "3"}

	assert.Equal(t, expected, appendNonDuplicates(testArray, testAppended))
}

func TestGetQueryUrls(t *testing.T) {
	// FLAKY
	getQueryUrls = actualGetQueryUrls
	urls := getQueryUrls("wikipedia")
	wikipediaFound := false
	for _, url := range urls {
		if strings.Contains(url, "wikipedia.org") {
			wikipediaFound = true
		}
	}
	assert.Equal(t, true, wikipediaFound)
}

func dummyGetQueryUrls(query string) []string {
	return []string{"fake://" + string(query[0])}
}

func TestGetTextUrls(t *testing.T) {
	getQueryUrls = dummyGetQueryUrls

	testString := ""
	for i := 0; i < 100; i++ {
		testString += fmt.Sprintf("%v ", i)
	}

	urls := getTextUrls(hooke.NewText(testString, &testLanguage))

	assert.Equal(t, 4, len(urls))
	assert.Equal(t, "fake://3", urls[1])
}

func TestGoogle(t *testing.T) {
	getQueryUrls = dummyGetQueryUrls
	downloadWebsite = dummyDownloadWebsite

	testText := hooke.NewText("test test test", &testLanguage)
	texts := Google(testText, &testLanguage)

	assert.Equal(t, 1, len(texts))
	assert.Equal(t, "fake://t", texts[0].Raw)
}
