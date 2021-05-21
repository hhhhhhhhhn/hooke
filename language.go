package main

type Language struct {
	Code           string
	WordCharacters []byte
	IsStopWord     func(word string) bool
	Stem           func(word string) string
}
