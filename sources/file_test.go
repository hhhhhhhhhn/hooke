package sources

import (
	"strings"
	"testing"

	"github.com/hhhhhhhhhn/hooke"
	"github.com/stretchr/testify/assert"
)

var testLanguage = hooke.Language{
	Stem: func(word string) string {
		return word + "+"
	},
	IsWordCharacter: func(chr rune) bool {
		return strings.Contains("teststop0123456789", string(chr))
	},
	IsStopWord: func(word string) bool {
		return word == "stop"
	},
}

func TestFile(t *testing.T) {
	testText := File("example.txt", &testLanguage)
	assert.Equal(t, []string{"test+", "test+"}, testText.ProcessedWords)
}
