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

func TestTextTransform_GetFromSubstrUntilTheEOL(t *testing.T) {
	type args struct {
		src    string
		substr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				src:    "Definição: Perspicacia\nTeste",
				substr: "Definição: ",
			},
			want:    "Perspicacia",
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args{
				src:    "Definição: Perspicacia",
				substr: "Definição: ",
			},
			want:    "Perspicacia",
			wantErr: false,
		},
		{
			name: "Case 3",
			args: args{
				src:    "Definição: Perspicacia",
				substr: "Definiço: ",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tt := &TextTransform{}
			got, err := tt.GetFromSubstrUntilTheEOL(tc.args.src, tc.args.substr)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetFromSubstrUntilTheEOL() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("GetFromSubstrUntilTheEOL() got = %v, want %v", got, tc.want)
			}
		})
	}
}
