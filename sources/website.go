package sources

import (
	"io/ioutil"
	"net/http"
	"regexp"

	h "github.com/hhhhhhhhhn/hooke"
)

func Website(url string, language *h.Language) *h.Text {
	html := downloadWebsite(url)
	text := parseHtml(html)
	return h.NewText(text, language)
}

func actualDownloadWebsite(url string) string {
	response, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

var downloadWebsite = actualDownloadWebsite // For testing

func parseHtml(html string) (text string) {
	text = removeStyleRegex.ReplaceAllString(html, "")
	text = removeSciptRegex.ReplaceAllString(text, "")
	text = removeTagsRegex.ReplaceAllString(text, "")
	return text
}

var removeStyleRegex = regexp.MustCompile(`(i?)<style([\s\S]*?)</style>`)
var removeSciptRegex = regexp.MustCompile(`(i?)<script([\s\S]*?)</script>`)
var removeTagsRegex = regexp.MustCompile(`(i?)<[^>]+>`)
