package sources

import (
	"context"
	"strings"
	"time"

	h "github.com/hhhhhhhhhn/hooke"
	googlesearch "github.com/rocketlaunchr/google-search"
)

const queryLength = 32
const sleepSeconds = 2

func Google(text *h.Text, language *h.Language) (texts []*h.Text) {
	urls := getTextUrls(text)
	for _, url := range urls {
		texts = append(texts, Website(url, language))
	}
	return texts
}

func getTextUrls(text *h.Text) (urls []string) {
	queries := divideIntoParts(text.NonStopWords, queryLength)
	for _, query := range queries {
		queryString := strings.Join(query, " ")
		urls = appendNonDuplicates(urls, getQueryUrls(queryString))
		time.Sleep(time.Second * sleepSeconds)
	}
	return urls
}

func actualGetQueryUrls(query string) (urls []string) {
	results, err := googlesearch.Search(context.TODO(), query)
	if err != nil {
		return urls
	}
	for _, result := range results {
		urls = append(urls, result.URL)
	}
	return urls
}

var getQueryUrls = actualGetQueryUrls

func appendNonDuplicates(original []string, appended []string) []string {
	for _, appendedElement := range appended {
		if !elementInArray(original, appendedElement) {
			original = append(original, appendedElement)
		}
	}
	return original
}

func elementInArray(array []string, appendedElement string) bool {
	for _, originalElement := range array {
		if originalElement == appendedElement {
			return true
		}
	}
	return false
}

func divideIntoParts(array []string, partLength int) [][]string {
	var output [][]string
	for i := 0; i <= len(array); i += partLength {
		output = append(output, array[i:min(i+partLength, len(array))])
	}
	return output
}

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}
