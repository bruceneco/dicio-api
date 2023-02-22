package lib

import (
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
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

func (tt *TextTransform) GetFromSubstrUntilTheEOL(src, substr string) (string, error) {
	if !strings.Contains(src, substr) {
		return "", fmt.Errorf("substr could not be found in src")
	}
	return strings.TrimSpace(strings.Split(strings.Split(src, substr)[1], "\n")[0]), nil
}
