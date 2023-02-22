package models

type Meaning struct {
	Type    string `json:"type,omitempty"`
	Meaning string `json:"meaning,omitempty"`
}

type Definition struct {
	GrammClass  string   `json:"grammClass,omitempty"`
	SyllabicSep []string `json:"syllabicSep,omitempty"`
	Plural      string   `json:"plural,omitempty"`
}
