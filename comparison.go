package hooke

type Comparison struct {
	Text1           *Text
	Text2           *Text
	Distance        int
	Matches         []Match
	Clusters        []Cluster
	SimilarityScore float64
}

// compares the two texts, returning a Comparison object
func NewComparison(text1 *Text, text2 *Text, distance int) *Comparison {
	comparison := &Comparison{
		Text1:    text1,
		Text2:    text2,
		Distance: distance,
	}
	comparison.compare()
	return comparison
}

// After the texts have been created, does the actual comparisons.
func (c *Comparison) compare() {
	c.findMatches()
	c.cluster()
	c.score()
}

// sets the matches field
func (c *Comparison) findMatches() {
	for text1Index, text1Word := range c.Text1.ProcessedWords {
		for text2Index, text2Word := range c.Text2.ProcessedWords {
			if text1Word == text2Word {
				c.Matches = append(c.Matches, Match{
					Text1Index: text1Index,
					Text2Index: text2Index,
				})
			}
		}
	}
}

// does the clustering
func (c *Comparison) cluster() {
	for i := range c.Matches {
		c.addMatchToClusters(&c.Matches[i])
	}
	c.getClusterStartsAndEnds()
	c.scoreClusters()
}

func (c *Comparison) addMatchToClusters(match *Match) {
	clusters := c.clustersMatchIsIn(match)

	if len(clusters) == 0 {
		c.Clusters = append(c.Clusters, Cluster{
			Matches: []*Match{match},
		})
		return
	}

	c.mergeClustersToFirstOne(clusters)
	clusters[0].Matches = append(clusters[0].Matches, match)

}

func (c *Comparison) clustersMatchIsIn(match *Match) []*Cluster {
	var inClusters []*Cluster

	for i := range c.Clusters {
		if c.Clusters[i].isWithinDistance(*match, c.Distance) {
			inClusters = append(inClusters, &c.Clusters[i])
		}
	}

	return inClusters
}

func (c *Comparison) mergeClustersToFirstOne(clusters []*Cluster) {
	for _, cluster := range clusters[1:] {
		clusters[0].Matches = append(clusters[0].Matches, cluster.Matches...)
		c.removeCluster(cluster)
	}
}

func (c *Comparison) removeCluster(cluster *Cluster) {
	for i := range c.Clusters {
		if &c.Clusters[i] == cluster {
			c.Clusters = append(c.Clusters[:i], c.Clusters[i+1:]...)
			return
		}
	}
}

func (c *Comparison) getClusterStartsAndEnds() {
	for i := range c.Clusters {
		c.Clusters[i].getStartAndEnd()
	}
}

func (c *Comparison) scoreClusters() {
	for i := range c.Clusters {
		c.Clusters[i].score()
	}
}

// sets the score based on the clusters
// read TestScoreCluster and TestScore in comparison_test.go
func (c *Comparison) score() {
	var scoreSum float64 = 0
	for _, cluster := range c.Clusters {
		scoreSum += float64(cluster.Score)
	}

	c.SimilarityScore = scoreSum /
		float64(len(c.Text1.ProcessedWords)+len(c.Text2.ProcessedWords))
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

func (cl Cluster) isWithinDistance(match Match, distance int) bool {
	for _, clusterMatch := range cl.Matches {
		if matchDistance(*clusterMatch, match) <= distance {
			return true
		}
	}
	return false
}

func matchDistance(match1 Match, match2 Match) int {
	text1Distance := abs(match1.Text1Index - match2.Text1Index)
	text2Distance := abs(match1.Text2Index - match2.Text2Index)

	if text1Distance >= text2Distance {
		return text1Distance

	}
	return text2Distance
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (c *Cluster) getStartAndEnd() {
	c.Text1Start = maxInt
	c.Text2Start = maxInt
	c.Text1End = -1
	c.Text2End = -1

	for _, match := range c.Matches {
		c.Text1Start = min(c.Text1Start, match.Text1Index)
		c.Text1End = max(c.Text1End, match.Text1Index)
		c.Text2Start = min(c.Text2Start, match.Text2Index)
		c.Text2End = max(c.Text2End, match.Text2Index)
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

const maxInt int = int(^uint(0) >> 1)

func (cl *Cluster) score() {
	cl.Score = len(cl.Matches) * len(cl.Matches) *
		(cl.Text1End - cl.Text1Start + cl.Text2End - cl.Text2Start)
}
