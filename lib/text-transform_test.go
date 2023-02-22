package lib

import (
	"testing"
)

func TestTextTransform_RemoveAccents(t *testing.T) {
	tt := NewTextTransform()

	tests := []struct {
		input    string
		expected string
	}{
		{"áàãâä", "aaaaa"},
		{"éèêë", "eeee"},
		{"íìîï", "iiii"},
		{"óòõôö", "ooooo"},
		{"úùûü", "uuuu"},
		{"ç", "c"},
		{"ÁÀÃÂÄ", "AAAAA"},
		{"ÉÈÊË", "EEEE"},
		{"ÍÌÎÏ", "IIII"},
		{"ÓÒÕÔÖ", "OOOOO"},
		{"ÚÙÛÜ", "UUUU"},
		{"Ç", "C"},
		{"hello world", "hello world"},
		{"1234", "1234"},
		{"", ""},
	}

	for _, test := range tests {
		result, err := tt.RemoveAccents(test.input)
		if err != nil {
			t.Errorf("RemoveAccents(%q) returned error: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("RemoveAccents(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
