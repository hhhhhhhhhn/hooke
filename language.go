package hooke

type Language struct {
	IsWordCharacter func(chr rune) bool
	IsStopWord      func(word string) bool
	Stem            func(word string) string
	PostProcess     func(text *Text)
}
