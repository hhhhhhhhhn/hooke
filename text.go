package main

type Text struct {
	Language              *Language
	Raw                   string
	Words                 []string
	WordsStartIndexes     []int
	WordsEndIndexes       []int
	ProcessedWords        []string
	ProcessedWordsIndexes []string
}

// takes a text with only raw and creates the other fields
func (t *Text) process() {
}

// divides words and wordIndexes based on languages acceptable words
func (t *Text) getWords() {
}

// sets processedWords and processedWordsIndex based on language stopwords
func (t *Text) removeStopWords() {
}

// applies the language's stem algorithm on the processedWords
func (t *Text) stemWords() {
}

// creates and processes string into a text object
func NewText(text string, language *Language) *Text {
	return &Text{}
}
