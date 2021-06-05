package sources

import (
	"io/ioutil"

	h "github.com/hhhhhhhhhn/hooke"
)

func File(filename string, language *h.Language) *h.Text {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		fileContents = []byte{}
	}
	return h.NewText(string(fileContents), language)
}
