package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLanguage = &Language{
	Code:           "test",
	WordCharacters: []byte{'t', 'e', 's', 't', 's', 't', 'o', 'p'},
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

func TestSplitWords(t *testing.T) {
	testText.getWords()
	assert.Equal(t, []string{"stop", "test"}, testText.Words, "splitting words")
	assert.Equal(t, []int{1, 7}, testText.WordsStartIndexes, "indexing word starts")
	assert.Equal(t, []int{4, 10}, testText.WordsEndIndexes, "indexing word end")
}

func TestRemoveStopWords(t *testing.T) {
	testText.removeStopWords()
	assert.Equal(t, []string{"test"}, testText.ProcessedWords, "remove stopwords")
	assert.Equal(t, []int{1}, testText.ProcessedWordsIndexes, "set processed word indexes")
}

func TestStemWords(t *testing.T) {
	testText.stemWords()
	assert.Equal(t, []string{"tes"}, testText.ProcessedWords, "stem words")
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
