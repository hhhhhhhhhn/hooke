package utils

import (
	"fmt"

	h "github.com/hhhhhhhhhn/hooke"
)

func PrintComparison(c *h.Comparison) (int, error) {
	return fmt.Print(formatComparison(c))
}

func formatComparison(c *h.Comparison) (formatted string) {
	for i, cluster := range c.Clusters {
		formatted +=
			"CLUSTER " + fmt.Sprint(i+1) + ":\n" +
				formatCluster(c, &cluster) + "\n\n\n"
	}
	return formatted[:len(formatted)-2] // strips last 2 newlines
}

func formatCluster(c *h.Comparison, cl *h.Cluster) string {
	return "TEXT 1:\n" +
		formatText(c.Text1, cl.Text1Start, cl.Text1End) +
		"\n\nTEXT 2:\n" +
		formatText(c.Text2, cl.Text2Start, cl.Text2End)
}

func formatText(t *h.Text, start int, end int) string {
	return processedIndexToRaw(t, start-5, start-1) +
		highlight +
		processedIndexToRaw(t, start, end) +
		reset +
		processedIndexToRaw(t, end+1, end+5)
}

func processedIndexToRaw(t *h.Text, start int, end int) string {
	start, end, inBounds := bound(start, end, len(t.ProcessedWordsIndexes))

	if !inBounds {
		return ""
	}

	startWord := t.ProcessedWordsIndexes[start]
	endWord := t.ProcessedWordsIndexes[end]

	return wordIndexToRaw(t, startWord, endWord)
}

func wordIndexToRaw(t *h.Text, start int, end int) string {
	startChar := t.WordsStartIndexes[start]

	if end == len(t.WordsStartIndexes)-1 { // is last word
		return t.Raw[startChar:]
	}

	endChar := t.WordsStartIndexes[end+1]
	return t.Raw[startChar:endChar]
}

func bound(start int, end int, length int) (newStart int, newEnd int, inBounds bool) {
	newStart = start
	newEnd = end
	if start < 0 {
		newStart = 0
	} else if start >= length {
		return 0, 0, false
	}
	if end >= length {
		newEnd = length - 1
	} else if end < 0 {
		return 0, 0, false
	}
	return newStart, newEnd, true
}

const highlight = "\033[7m"
const reset = "\033[0m"
