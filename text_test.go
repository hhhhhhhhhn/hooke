package hooke

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLanguage = &Language{
	IsWordCharacter: func(chr rune) bool {
		return strings.Contains("teststop", string(chr))
	},
	IsStopWord: func(word string) bool {
		return word == "stop"
	},
	Stem: func(word string) string {
		return word[:len(word)-1]
	},
}
var testText = &Text{
	Raw:      " Stop  Test zz",
	Language: testLanguage,
}

func TestWordGetStartAndEndIndexes(t *testing.T) {
	testTextCopy := *testText
	testTextCopy.getWordStartsAndEnds()

	assert.Equal(t, []int{1, 7}, testTextCopy.WordsStartIndexes)
	assert.Equal(t, []int{5, 11}, testTextCopy.WordsEndIndexes)
}
func TestSplitWords(t *testing.T) {
	testText.getWords()
	assert.Equal(t, []string{"stop", "test"}, testText.Words, "splitting words")
	assert.Equal(t, []int{1, 7}, testText.WordsStartIndexes, "indexing word starts")
	assert.Equal(t, []int{5, 11}, testText.WordsEndIndexes, "indexing word end")
}

func TestGetWordsWithoutEnding(t *testing.T) {
	text := &Text{
		Raw:      "stop test test",
		Language: testLanguage,
	}

	text.getWords()
	assert.Equal(t, 3, len(text.Words))
	assert.Equal(t, []int{0, 5, 10}, text.WordsStartIndexes)
	assert.Equal(t, []string{"stop", "test", "test"}, text.Words)
}

func TestRemoveStopWords(t *testing.T) {
	testText.removeStopWords()
	assert.Equal(t, []string{"test"}, testText.NonStopWords, "remove stopwords")
	assert.Equal(t, []int{1}, testText.NonStopWordsIndexes)
}

func TestStemWords(t *testing.T) {
	testText.stemWords()
	assert.Equal(t, []string{"tes"}, testText.ProcessedWords, "stem words")
	assert.Equal(t, []int{1}, testText.ProcessedWordsIndexes)
	testText.NonStopWordsIndexes[0] = -1 // test they are different arrays
	assert.Equal(t, []int{1}, testText.ProcessedWordsIndexes)
}

func TestProcess(t *testing.T) {
	testText := &Text{
		Raw:      "stop test test",
		Language: testLanguage,
	}
	testText.process()
	assert.Equal(t, []string{"tes", "tes"}, testText.ProcessedWords, "process does everything")
}

func TestNewText(t *testing.T) {
	testText := NewText("ttt stop test test", testLanguage)
	assert.Equal(t, []string{"tt", "tes", "tes"}, testText.ProcessedWords, "NewText processes")
}

func TestPostProcess(t *testing.T) {
	test := *testLanguage
	test.PostProcess = func(t *Text) {
		for i := range t.ProcessedWords {
			t.ProcessedWords[i] += "a"
		}
	}
	testText := NewText("ttt stop test test", &test)
	assert.Equal(t, []string{"tta", "tesa", "tesa"}, testText.ProcessedWords, "NewText processes")
}

func TestEmptyNewText(t *testing.T) {
	testText := NewText("", &Language{})
	assert.Equal(t, 0, len(testText.ProcessedWords))
}
