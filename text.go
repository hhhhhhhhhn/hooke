package main

import "strings"

type Text struct {
	Language              *Language
	Raw                   string
	LowercaseRaw          string
	Words                 []string
	WordsStartIndexes     []int
	WordsEndIndexes       []int
	ProcessedWords        []string
	ProcessedWordsIndexes []int
}

// creates and processes string into a text object
func NewText(text string, language *Language) *Text {
	textObject := &Text{
		Raw:      text,
		Language: language,
	}
	textObject.process()
	return textObject
}

// takes a text with only raw and creates the other fields
func (t *Text) process() {
	t.getWords()
	t.removeStopWords()
	t.stemWords()
	t.postProcess()
}

// divides words and wordIndexes based on languages acceptable words + lowercase
func (t *Text) getWords() {
	t.getWordStartsAndEnds()

	wordCount := len(t.WordsStartIndexes)
	t.Words = make([]string, wordCount)

	for i := 0; i < wordCount; i++ {
		t.Words[i] = t.LowercaseRaw[t.WordsStartIndexes[i]:t.WordsEndIndexes[i]]
	}
}

func (t *Text) getWordStartsAndEnds() {
	wordCount := t.countWords()

	t.WordsStartIndexes = make([]int, wordCount)
	t.WordsEndIndexes = make([]int, wordCount)

	inWord := false
	word := 0

	for i, chr := range t.LowercaseRaw {
		if t.Language.isWordCharacter(chr) {
			if !inWord {
				inWord = true
				t.WordsStartIndexes[word] = i
			}
		} else { // chr is not a word character
			if inWord {
				inWord = false
				t.WordsEndIndexes[word] = i
				word++
			}
		}
	}
	if inWord {
		t.WordsEndIndexes[word] = len(t.LowercaseRaw)
	}
}

func (t *Text) countWords() (count int) {
	t.getLowercase()
	count = 0
	inWord := false

	for _, chr := range t.LowercaseRaw {
		if t.Language.isWordCharacter(chr) {
			if !inWord {
				inWord = true
				count++
			}
		} else { // chr is not a word character
			if inWord {
				inWord = false
			}
		}
	}
	return count
}

func (t *Text) getLowercase() {
	t.LowercaseRaw = strings.ToLower(t.Raw)
}

// sets processedWords and processedWordsIndex based on language stopwords
func (t *Text) removeStopWords() {
	t.ProcessedWords = make([]string, len(t.Words)) // length is upper limit
	t.ProcessedWordsIndexes = make([]int, len(t.Words))
	processedWord := 0

	for i, word := range t.Words {
		if !t.Language.IsStopWord(word) {
			t.ProcessedWords[processedWord] = word
			t.ProcessedWordsIndexes[processedWord] = i
			processedWord++
		}
	}

	t.ProcessedWords = t.ProcessedWords[:processedWord] // sets actual length
	t.ProcessedWordsIndexes = t.ProcessedWordsIndexes[:processedWord]
}

// applies the language's stem algorithm on the processedWords
func (t *Text) stemWords() {
	for i := range t.ProcessedWords {
		t.ProcessedWords[i] = t.Language.Stem(t.ProcessedWords[i])
	}
}

func (t *Text) postProcess() {
	if t.Language.PostProcess != nil {
		t.Language.PostProcess(t)
	}
}
