package lib

import "github.com/gocolly/colly"

type Scrap struct{}

func NewScrap() *Scrap {
	return &Scrap{}
}

func (s *Scrap) GetColl() *colly.Collector {
	return colly.NewCollector(colly.AllowURLRevisit())
}
