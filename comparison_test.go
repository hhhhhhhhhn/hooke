package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMatches(t *testing.T) {
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

func TestAbs(t *testing.T) {
	assert.Equal(t, 2, abs(-2))
	assert.Equal(t, 2, abs(2))
	assert.Equal(t, 0, abs(-0))
}

func TestMatchDistance(t *testing.T) {
	// NOTE: Its Chebyshev distance
	match1 := Match{
		Text1Index: 0,
		Text2Index: 1,
	}
	match2 := Match{
		Text1Index: -100,
		Text2Index: 50,
	}
	assert.Equal(t, 100, matchDistance(match1, match2))

	match1 = Match{
		Text1Index: 1,
		Text2Index: 0,
	}
	match2 = Match{
		Text1Index: 50,
		Text2Index: -100,
	}

	assert.Equal(t, 100, matchDistance(match2, match1))
}

func TestClusterIsWithinDistance(t *testing.T) {
	match := Match{
		Text1Index: 50,
		Text2Index: 50,
	}
	testCluster := Cluster{
		Matches: []*Match{
			{
				Text1Index: 0,
				Text2Index: 0,
			},
			{
				Text1Index: 100,
				Text2Index: 100,
			},
		},
	}
	assert.Equal(t, false, testCluster.isWithinDistance(match, 10))
	assert.Equal(t, true, testCluster.isWithinDistance(match, 50))
}

func TestMatchInClusters(t *testing.T) {
	testComparison := Comparison{
		Clusters: []Cluster{
			{
				Matches: []*Match{
					{
						Text1Index: 100,
						Text2Index: 100,
					},
				},
			},
			{
				Matches: []*Match{
					{
						Text1Index: 100,
						Text2Index: 100,
					},
					{
						Text1Index: 0,
						Text2Index: 0,
					},
				},
			},
		},
		Distance: 10,
	}
	testMatch := Match{
		Text1Index: 1,
		Text2Index: 1,
	}
	expected := []*Cluster{
		&testComparison.Clusters[1],
	}
	assert.Equal(t, expected, testComparison.clustersMatchIsIn(&testMatch))
}

func TestAddMatchToClusters(t *testing.T) {
	testComparison := Comparison{
		Clusters: []Cluster{
			{
				Matches: []*Match{
					{
						Text1Index: 100,
						Text2Index: 100,
					},
				},
			},
			{
				Matches: []*Match{
					{
						Text1Index: 100,
						Text2Index: 100,
					},
					{
						Text1Index: 0,
						Text2Index: 0,
					},
				},
			},
		},
		Distance: 10,
	}
	testMatch := Match{
		Text1Index: 1,
		Text2Index: 1,
	}

	testComparison.addMatchToClusters(&testMatch)

	expected := Cluster{
		Matches: []*Match{
			testComparison.Clusters[1].Matches[0],
			testComparison.Clusters[1].Matches[1],
			&testMatch,
		},
	}
	assert.Equal(t, expected, testComparison.Clusters[1])
}

func TestRemoveCluster(t *testing.T) {
	testComparison := Comparison{
		Clusters: []Cluster{
			{Matches: []*Match{{0, 0}}},
			{Matches: []*Match{{1, 1}}},
			{Matches: []*Match{{2, 2}}},
		},
	}
	removed := &testComparison.Clusters[1]
	expected := []Cluster{
		testComparison.Clusters[0],
		testComparison.Clusters[2],
	}
	testComparison.removeCluster(removed)

	assert.Equal(t, expected, testComparison.Clusters)
}

func TestMergeClusters(t *testing.T) {
	testComparison := Comparison{
		Clusters: []Cluster{
			{
				Matches: []*Match{{0, 0}, {1, 1}},
			},
			{
				Matches: []*Match{{2, 2}, {3, 3}},
			},
			{
				Matches: []*Match{{4, 4}, {5, 5}},
			},
		},
	}

	expectedMatches := []*Match{
		testComparison.Clusters[1].Matches[0],
		testComparison.Clusters[1].Matches[1],
		testComparison.Clusters[2].Matches[0],
		testComparison.Clusters[2].Matches[1],
	}

	testComparison.mergeClustersToFirstOne([]*Cluster{
		&testComparison.Clusters[1],
		&testComparison.Clusters[2],
	})

	assert.Equal(t, 2, len(testComparison.Clusters))
	assert.Equal(t, expectedMatches, testComparison.Clusters[1].Matches)
}

func TestGetClusterStartAndEnd(t *testing.T) {
	testCluster := Cluster{
		Matches: []*Match{
			{4, 10},
			{15, 2},
			{5, 5},
		},
	}
	testCluster.getStartAndEnd()

	assert.Equal(t, 4, testCluster.Text1Start)
	assert.Equal(t, 15, testCluster.Text1End)
	assert.Equal(t, 2, testCluster.Text2Start)
	assert.Equal(t, 10, testCluster.Text2End)
}

func TestGetClusterStartsAndEnds(t *testing.T) {
	testComparison := Comparison{
		Clusters: []Cluster{
			{
				Matches: []*Match{{0, 0}, {1, 1}},
			},
			{
				Matches: []*Match{{4, 4}, {5, 5}},
			},
			{
				Matches: []*Match{{2, 2}, {3, 3}},
			},
		},
	}
	testComparison.getClusterStartsAndEnds()

	assert.Equal(t, 4, testComparison.Clusters[1].Text1Start)
	assert.Equal(t, 4, testComparison.Clusters[1].Text2Start)
	assert.Equal(t, 5, testComparison.Clusters[1].Text1End)
	assert.Equal(t, 5, testComparison.Clusters[1].Text2End)
}

func TestScoreCluster(t *testing.T) {
	testCluster := Cluster{
		Matches:    []*Match{{}, {}, {}},
		Text1Start: 5,
		Text1End:   10,
		Text2Start: 20,
		Text2End:   30,
	}
	testCluster.score()
	// Score should be equal to the length * density^2 of the cluster,
	// where the length is the sum of the length of each text,
	// and the density is the amount of matches
	//
	// i.e. (Nº Matches)^2 * (Text1End - Text1Start + Text2End - Text2End)
	//
	// substituting for this case
	// 3^2 * (10 - 5 + 30 - 20)
	// = 135
	assert.Equal(t, 135, testCluster.Score)
}

func TestScoreClusters(t *testing.T) {
	testComparison := Comparison{
		Clusters: []Cluster{
			{
				Matches:    []*Match{{}, {}, {}},
				Text1Start: 0,
				Text2Start: 0,
				Text1End:   5,
				Text2End:   10,
			},
			{
				Matches:    []*Match{{}, {}, {}},
				Text1Start: 0,
				Text2Start: 0,
				Text1End:   5,
				Text2End:   5,
			},
		},
	}
	testComparison.scoreClusters()
	// substituting for this case:
	// 3^2 * (5 - 0 + 10 - 0)
	// = 135
	// 3^2 * (5 - 0 + 5 - 0)
	// = 90
	assert.Equal(t, 135, testComparison.Clusters[0].Score)
	assert.Equal(t, 90, testComparison.Clusters[1].Score)

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
	assert.Greater(t, testComparison.SimilarityScore, float64(0))

	assert.Equal(t, 2, len(testComparison.Matches))
	assert.Equal(t, 1, len(testComparison.Clusters))
}

func TestNewComparison(t *testing.T) {
	testComparison := NewComparison(
		&Text{ProcessedWords: []string{"a", "b", "c"}},
		&Text{ProcessedWords: []string{"0", "b", "c", "3", "4"}},
		2,
	)

	assert.Greater(t, testComparison.SimilarityScore, float64(0))

	assert.Equal(t, 2, len(testComparison.Matches))
	assert.Equal(t, 1, len(testComparison.Clusters))
}

func TestEmptyNewComparison(t *testing.T) {
	testComparison := NewComparison(&Text{}, &Text{}, 0)

	assert.Equal(t, 0, len(testComparison.Matches))
	assert.Equal(t, 0, len(testComparison.Clusters))
}
