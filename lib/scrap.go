package lib

import "github.com/gocolly/colly"

type Scrap struct {
	*colly.Collector
}

func NewScrap() Scrap {
	return Scrap{Collector: colly.NewCollector()}
}
