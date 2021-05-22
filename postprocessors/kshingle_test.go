package postprocessors

import (
	"strings"
	"testing"

	"github.com/hhhhhhhhhn/hooke"
	"github.com/stretchr/testify/assert"
)

var testLanguage = hooke.Language{
	Code: "test",
	Stem: func(word string) string {
		return word + "+"
	},
	IsWordCharacter: func(chr rune) bool {
		return strings.Contains("teststop", string(chr))
	},
	IsStopWord: func(word string) bool {
		return word == "stop"
	},
}

func TestShingle(t *testing.T) {
	testWords := []string{"0", "1", "2", "3", "4"}
	expected := []string{"0 1 2", "1 2 3", "2 3 4"}
	assert.Equal(t, expected, shingle(testWords, 3))
}

func TestKShingle(t *testing.T) {
	languageCopy := testLanguage
	languageCopy.PostProcess = KShingle(2)

	testText := hooke.NewText("t e s t", &languageCopy)

	expectedWords := []string{"t+ e+", "e+ s+", "s+ t+"}
	expectedIndexes := []int{0, 1, 2}

	assert.Equal(t, expectedWords, testText.ProcessedWords)
	assert.Equal(t, expectedIndexes, testText.ProcessedWordsIndexes)
}
