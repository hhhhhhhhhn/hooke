package main

import (
	h "github.com/hhhhhhhhhn/hooke"
	"github.com/hhhhhhhhhn/hooke/languages"
	"github.com/hhhhhhhhhn/hooke/sources"
	"github.com/hhhhhhhhhn/hooke/utils"
	"github.com/hhhhhhhhhn/hooke/postprocessors"

	"io/ioutil"
	"os"
	"sort"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)

	language := &languages.Engish
	language.PostProcess = postprocessors.KShingle(2)

	text := h.NewText(string(data), &languages.Engish)
	texts := sources.Google(text, &languages.Engish)


	comparison := h.NewComparison(text, texts[0], 3)
	sort.Slice(comparison.Clusters, func(i, j int) bool {return len(comparison.Clusters[i].Matches) > len(comparison.Clusters[j].Matches) })

	utils.PrintComparison(comparison, 1000)
}
