package main

type Comparison struct {
	Text1           *Text
	Text2           *Text
	Distance        int
	Matches         []Match
	Clusters        []Cluster
	SimilarityScore float64
}

// After the texts have been created, does the actual comparisons.
func (c *Comparison) compare() {
}

// sets the matches field
func (c *Comparison) findMatches() {
}

// does the clustering
func (c *Comparison) cluster() {
}

// sets the score based on the clusters
// read end of TestCluster and TestScore in comparison_test.go
func (c *Comparison) score() {
}

// compares the two texts, returning a Comparison object
func NewComparison(text1 *Text, text2 *Text, distance int) *Comparison {
	return &Comparison{}
}

type Match struct {
	Text1Index int
	Text2Index int
}

type Cluster struct {
	Matches    []*Match
	Text1Start int
	Text1End   int
	Text2Start int
	Text2End   int
	Score      int
}
