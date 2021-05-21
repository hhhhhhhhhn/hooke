package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMatchs(t *testing.T) {
	testComparison := &Comparison{
		Text1: &Text{ProcessedWords: []string{"one", "two", "three"}},
		Text2: &Text{ProcessedWords: []string{"hello", "there", "two", "three", "bre"}},
	}
	testComparison.findMatches()
	expected := []Match{
		{ // "two"
			Text1Index: 1,
			Text2Index: 2,
		},
		{ // "three"
			Text1Index: 2,
			Text2Index: 3,
		},
	}
	assert.Equal(t, expected, testComparison.Matches)
}

func TestCluster(t *testing.T) {
	testComparison := &Comparison{
		Matches: []Match{
			{
				Text1Index: 20,
				Text2Index: 20,
			},
			{
				Text1Index: 0,
				Text2Index: 1,
			},
			{
				Text1Index: 2,
				Text2Index: 3,
			},
		},
		Distance: 5,
	}
	testComparison.cluster()
	expectedMatches := []*Match{
		&(testComparison.Matches[1]),
		&(testComparison.Matches[2]),
	}

	assert.Equal(t, 2, len(testComparison.Clusters))

	assert.Equal(t, 1, len(testComparison.Clusters[0].Matches))

	assert.Equal(t, expectedMatches, testComparison.Clusters[1].Matches)

	assert.Equal(t, 0, testComparison.Clusters[1].Text1Start)
	assert.Equal(t, 2, testComparison.Clusters[1].Text1End)
	assert.Equal(t, 1, testComparison.Clusters[1].Text2Start)
	assert.Equal(t, 3, testComparison.Clusters[1].Text2End)

	// Score should be equal to the length * density^2 of the cluster,
	// where the length is the sum of the length of each text,
	// and the density is the amount of matches
	//
	// i.e. (Nº Matches)^2 * (Text1End - Text1Start + Text2End - Text2End)
	//
	// substituting for this case:
	// (2)^2 * (2 - 0 + 3 - 1)
	// = 16

	assert.Equal(t, 16, testComparison.Clusters[1].Score)
}

func TestScore(t *testing.T) {
	testComparison := &Comparison{
		Clusters: []Cluster{
			{
				Score: 10,
			},
			{
				Score: 20,
			},
		},
		Text1: &Text{
			ProcessedWords: []string{"1", "2", "3", "4"},
		},
		Text2: &Text{
			ProcessedWords: []string{"1", "2"},
		},
	}
	testComparison.score()

	// The similarity index should be the sum of all cluster scores / length,
	// where the length is the sum of the lengths of both texts
	//
	// i.e. (Sum of cluster scores) / (len(Text1.ProcessedWords) + len(Text2.ProcessedWords))
	//
	// substituting for this case
	// (30) / (4 + 2)
	// = 5

	assert.InDelta(t, 5, testComparison.SimilarityScore, 0.1)
}

func TestCompare(t *testing.T) {
	testComparison := &Comparison{
		Text1: &Text{ProcessedWords: []string{"a", "b", "c"}},
		Text2: &Text{ProcessedWords: []string{"0", "1", "2", "3", "4"}},
	}
	testComparison.compare()
	assert.InDelta(t, 0, testComparison.SimilarityScore, 0.1)

	testComparison = &Comparison{
		Text1:    &Text{ProcessedWords: []string{"a", "b", "c"}},
		Text2:    &Text{ProcessedWords: []string{"0", "b", "c", "3", "4"}},
		Distance: 2,
	}
	testComparison.compare()
	assert.Greater(t, testComparison.SimilarityScore, 0)

	assert.Equal(t, 2, len(testComparison.Matches))
	assert.Equal(t, 1, len(testComparison.Clusters))
}

func TestNewComparison(t *testing.T) {
	testComparison := NewComparison(
		&Text{ProcessedWords: []string{"a", "b", "c"}},
		&Text{ProcessedWords: []string{"0", "b", "c", "3", "4"}},
		2,
	)

	assert.Greater(t, testComparison.SimilarityScore, 0)

	assert.Equal(t, 2, len(testComparison.Matches))
	assert.Equal(t, 1, len(testComparison.Clusters))
}

func TestEmptyNewComparison(t *testing.T) {
	testComparison := NewComparison(&Text{}, &Text{}, 0)

	assert.Equal(t, 0, len(testComparison.Matches))
	assert.Equal(t, 0, len(testComparison.Clusters))
}
