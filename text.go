package hooke

import "strings"

type Text struct {
	Language            *Language
	Raw                 string
	LowercaseRaw        string
	Words               []string
	WordsStartIndexes   []int
	WordsEndIndexes     []int
	NonStopWords        []string
	NonStopWordsIndexes []int
	ProcessedWords      []string
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

	for i := 0; i < len(t.WordsStartIndexes); i++ {
		t.Words = append(
			t.Words,
			t.LowercaseRaw[t.WordsStartIndexes[i]:t.WordsEndIndexes[i]],
		)
	}
}

func (t *Text) getWordStartsAndEnds() {
	t.getLowercase()
	inWord := false

	for i, chr := range t.LowercaseRaw {
		if t.Language.isWordCharacter(chr) {
			if !inWord {
				inWord = true
				t.WordsStartIndexes = append(t.WordsStartIndexes, i)
			}
		} else { // chr is not a word character
			if inWord {
				inWord = false
				t.WordsEndIndexes = append(t.WordsEndIndexes, i)
			}
		}
	}
	if inWord {
		t.WordsEndIndexes = append(t.WordsEndIndexes, len(t.LowercaseRaw))
	}
}

func (t *Text) getLowercase() {
	t.LowercaseRaw = strings.ToLower(t.Raw)
}

// sets NonStopWords and processedWordsIndex based on language stopwords
func (t *Text) removeStopWords() {
	for i, word := range t.Words {
		if !t.Language.IsStopWord(word) {
			t.NonStopWords = append(t.NonStopWords, word)
			t.NonStopWordsIndexes = append(t.NonStopWordsIndexes, i)
		}
	}
}

// applies the language's stem algorithm on the processedWords
func (t *Text) stemWords() {
	for _, word := range t.NonStopWords {
		t.ProcessedWords = append(t.ProcessedWords, t.Language.Stem(word))
	}
}

func (t *Text) postProcess() {
	if t.Language.PostProcess != nil {
		t.Language.PostProcess(t)
	}
}
