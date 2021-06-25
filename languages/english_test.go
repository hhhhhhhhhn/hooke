package languages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnglish(t *testing.T) {
	assert.Equal(t, true, Engish.IsStopWord("the"))

	assert.Equal(t, false, Engish.IsStopWord("unusual"))

	assert.Equal(t, "jazzi", Engish.Stem("jazzy"))

	assert.Equal(t, true, Engish.IsWordCharacter('a'))
	assert.Equal(t, false, Engish.IsWordCharacter('ñ'))
}
