package main

type Language struct {
	Code            string
	isWordCharacter func(chr rune) bool
	IsStopWord      func(word string) bool
	Stem            func(word string) string
	PostProcess     func(text *Text)
}
