package languages

import (
	"strings"

	"github.com/bbalet/stopwords"
	h "github.com/hhhhhhhhhn/hooke"
	"github.com/tebeka/snowball"
)

var englishStemmer *snowball.Stemmer
var englishStemmerInitalized = false

var Engish = h.Language{
	IsStopWord: func(word string) bool {
		return stopwords.CleanString(word, "en", true) != word+" "
	},
	Stem: func(word string) string {
		if !englishStemmerInitalized {
			englishStemmer, _ = snowball.New("english")
		}
		return englishStemmer.Stem(word)
	},
	IsWordCharacter: func(chr rune) bool {
		return strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", chr)
	},
}
