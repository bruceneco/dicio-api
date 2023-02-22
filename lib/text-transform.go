package lib

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type TextTransform struct{}

func NewTextTransform() *TextTransform {
	return &TextTransform{}
}

func (tt *TextTransform) RemoveAccents(s string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	res, _, err := transform.String(t, s)
	return res, err
}
