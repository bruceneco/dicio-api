package services

import (
	"fmt"
	"github.com/bruceneco/dicio-api/lib"
	"github.com/gocolly/colly"
	"math"
	"strings"
)

const (
	dicioURL = "https://www.dicio.com.br"
)

type ScrapService struct {
	logger *lib.Logger
	scrap  *lib.Scrap
}

func NewScrapService(logger *lib.Logger, scrap *lib.Scrap) *ScrapService {
	return &ScrapService{
		logger: logger,
		scrap:  scrap,
	}
}

func (s ScrapService) TopWords(nWords int) ([]string, error) {
	c := s.scrap.GetColl()
	nWords = int(math.Min(float64(nWords), float64(5000)))
	var words []string
	c.OnHTML(".list > li", func(e *colly.HTMLElement) {
		if len(words) == nWords {
			return
		}
		words = append(words, strings.TrimSpace(e.Text))
	})
	for i := 0.; i < math.Ceil(float64(nWords)/100.); i++ {
		err := c.Visit(fmt.Sprintf("%s/palavras-mais-buscadas/%g", dicioURL, i))
		if err != nil {
			s.logger.Errorf("can't access dicio site: %s", err.Error())
			return nil, fmt.Errorf("Não foi possível acessar o site do Dicio.")
		}
	}
	return words, nil
}
