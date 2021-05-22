package postprocessors

import (
	"testing"

	"github.com/hhhhhhhhhn/hooke"
	"github.com/stretchr/testify/assert"
)

func TestKShingle(t *testing.T) {
	testProcessor := KShingle(1)
	testComparison := hooke.NewComparison(
		hooke.NewText(""),
		hooke.NewText(""),
		10
	)
	assert.Equal(t, 1, 1)
}
