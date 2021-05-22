package postprocessors

import (
	"strings"

	"github.com/hhhhhhhhhn/hooke"
)

func KShingle(k int) func(t *hooke.Text) {
	return func(t *hooke.Text) {
		t.ProcessedWords = shingle(t.ProcessedWords, k)
		t.ProcessedWordsIndexes =
			t.ProcessedWordsIndexes[:len(t.ProcessedWordsIndexes)+1-k]
	}
}

func shingle(words []string, k int) (shingled []string) {
	for i := 0; i < len(words)+1-k; i++ {
		shingled = append(shingled, strings.Join(words[i:i+k], " "))
	}
	return shingled
}
