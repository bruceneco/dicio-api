package models

import "time"

type Meaning struct {
	Type    string `json:"type,omitempty"`
	Meaning string `json:"meaning,omitempty"`
}

type Definition struct {
	GrammClass  string   `json:"grammClass,omitempty"`
	SyllabicSep []string `json:"syllabicSep,omitempty"`
	Plural      string   `json:"plural,omitempty"`
}

type Example struct {
	Content string    `json:"content,omitempty"`
	Author  string    `json:"author,omitempty"`
	Date    time.Time `json:"date"`
}

type Citation struct {
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}
