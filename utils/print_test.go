package utils

import (
	"testing"

	"github.com/hhhhhhhhhn/hooke"
	"github.com/stretchr/testify/assert"
)

var testLanguage = hooke.Language{
	IsWordCharacter: func(chr rune) bool {
		return chr != ' '
	},
	IsStopWord: func(word string) bool { return false },
	Stem:       func(word string) string { return word },
}

func TestBound(t *testing.T) {
	start, end, inBounds := bound(5, 10, 10)
	assert.Equal(t, 5, start)
	assert.Equal(t, 9, end)
	assert.Equal(t, true, inBounds)

	start, end, inBounds = bound(-10, 5, 10)
	assert.Equal(t, 0, start)
	assert.Equal(t, 5, end)
	assert.Equal(t, true, inBounds)

	_, _, inBounds = bound(11, 10, 10)
	assert.Equal(t, false, inBounds)

	_, _, inBounds = bound(0, -5, 10)
	assert.Equal(t, false, inBounds)
}

func TestWordIndexToRaw(t *testing.T) {
	testText := hooke.Text{
		Raw:               "00 11 22 33 44 55",
		WordsStartIndexes: []int{0, 3, 6, 9, 12, 15},
		WordsEndIndexes:   []int{2, 5, 8, 11, 14, 17},
	}
	assert.Equal(t, "22 33 44 ", wordIndexToRaw(&testText, 2, 4))
}

func TestProcessedIndexToRaw(t *testing.T) {
	testText := hooke.Text{
		Raw:                   "00 11 22 33 44 55",
		WordsStartIndexes:     []int{0, 3, 6, 9, 12, 15},
		WordsEndIndexes:       []int{2, 5, 8, 11, 14, 17},
		ProcessedWordsIndexes: []int{1, 3, 4},
	}
	assert.Equal(t, "33 44 ", processedIndexToRaw(&testText, 1, 10))
}

func TestFormatText(t *testing.T) {
	testText1 := hooke.NewText("0 1 2 3 4 5 6 7 8 9", &testLanguage)
	testText2 := hooke.NewText("1 2 3 d e f g 7 i 9", &testLanguage)

	testComparison := hooke.NewComparison(testText1, testText2, 3)
	cluster := testComparison.Clusters[0]

	expectedText1 := "0 " + highlight + "1 2 3 " + reset + "4 5 6 7 8 "
	expectedText2 := highlight + "1 2 3 " + reset + "d e f g 7 "

	assert.Equal(t, expectedText1, formatText(testText1, cluster.Text1Start, cluster.Text1End))
	assert.Equal(t, expectedText2, formatText(testText2, cluster.Text2Start, cluster.Text2End))
}

func TestFormatCluster(t *testing.T) {
	testText1 := hooke.NewText("0 1 2 3 4 5 6 7 8 9", &testLanguage)
	testText2 := hooke.NewText("1 2 3 d e f g 7 i 9", &testLanguage)

	testComparison := hooke.NewComparison(testText1, testText2, 3)

	expected := "TEXT 1:\n" +
		"0 " + highlight + "1 2 3 " + reset + "4 5 6 7 8 \n" +
		"\n" +
		"TEXT 2:\n" +
		highlight + "1 2 3 " + reset + "d e f g 7 "

	assert.Equal(t, expected, formatCluster(testComparison, &testComparison.Clusters[0]))
}

func TestFormatComparison(t *testing.T) {
	testText1 := hooke.NewText("0 1 2 3 4 5 6 7 8 9", &testLanguage)
	testText2 := hooke.NewText("1 2 3 d e f g 7 i 9", &testLanguage)

	testComparison := hooke.NewComparison(testText1, testText2, 3)

	expected :=
		"CLUSTER 1:\n" +
			"TEXT 1:\n" +
			"0 " + highlight + "1 2 3 " + reset + "4 5 6 7 8 \n" +
			"\n" +
			"TEXT 2:\n" +
			highlight + "1 2 3 " + reset + "d e f g 7 \n" +
			"\n" +
			"\n" +
			"CLUSTER 2:\n" +
			"TEXT 1:\n" +
			"2 3 4 5 6 " + highlight + "7 8 9" + reset + "\n" +
			"\n" +
			"TEXT 2:\n" +
			"3 d e f g " + highlight + "7 i 9" + reset + "\n"

	assert.Equal(t, expected, formatComparison(testComparison))
}

func TestPrint(t *testing.T) {
	testText1 := hooke.NewText("0 1 2 3 4 5 6 7 8 9", &testLanguage)
	testText2 := hooke.NewText("1 2 3 d e f g 7 i 9", &testLanguage)

	testComparison := hooke.NewComparison(testText1, testText2, 3)

	expected :=
		"CLUSTER 1:\n" +
			"TEXT 1:\n" +
			"0 " + highlight + "1 2 3 " + reset + "4 5 6 7 8 \n" +
			"\n" +
			"TEXT 2:\n" +
			highlight + "1 2 3 " + reset + "d e f g 7 \n" +
			"\n" +
			"\n" +
			"CLUSTER 2:\n" +
			"TEXT 1:\n" +
			"2 3 4 5 6 " + highlight + "7 8 9" + reset + "\n" +
			"\n" +
			"TEXT 2:\n" +
			"3 d e f g " + highlight + "7 i 9" + reset + "\n"

	written, err := PrintComparison(testComparison)
	assert.Equal(t, len(expected), written)
	assert.Nil(t, err)
}
