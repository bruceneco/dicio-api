package services

import (
	"fmt"
	"github.com/bruceneco/dicio-api/lib"
	"github.com/bruceneco/dicio-api/models"
	"github.com/gocolly/colly"
	"math"
	"strings"
)

const (
	dicioURL = "https://www.dicio.com.br"
)

type ScrapService struct {
	logger        *lib.Logger
	scrap         *lib.Scrap
	textTransform *lib.TextTransform
}

func NewScrapService(logger *lib.Logger, scrap *lib.Scrap) *ScrapService {
	return &ScrapService{
		logger: logger,
		scrap:  scrap,
	}
}

func (s *ScrapService) TopWords(nWords int) ([]string, error) {
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
			s.logger.Warnf("can't access dicio site: %s", err.Error())
			return nil, fmt.Errorf("Não foi possível acessar o site do Dicio.")
		}
	}
	return words, nil
}

func (s *ScrapService) Meanings(word string) ([]*models.Meaning, error) {
	word, err := s.textTransform.RemoveAccents(strings.ToLower(word))
	if err != nil {
		s.logger.Warnf("can't remove accents from word: %s", err.Error())
		return nil, fmt.Errorf("Não foi possível remover os acentos da palavra para uma busca precisa.")
	}
	var meanings []*models.Meaning

	c := s.scrap.GetColl()
	c.OnHTML(".significado > span:not(.cl):not(.etim)", func(element *colly.HTMLElement) {
		meaning := models.Meaning{}
		kindMeanSplit := strings.SplitAfter(element.Text, "]")
		if len(kindMeanSplit) == 2 {
			meaning.Type = kindMeanSplit[0]
			meaning.Type = strings.TrimLeft(meaning.Type, "[")
			meaning.Type = strings.TrimRight(meaning.Type, "]")
			meaning.Type = strings.TrimSpace(meaning.Type)
			meaning.Meaning = kindMeanSplit[1]
		} else {
			meaning.Meaning = kindMeanSplit[0]
			meaning.Type = "Comum"
		}
		meaning.Meaning = strings.TrimSpace(meaning.Meaning)
		meanings = append(meanings, &meaning)
	})

	err = c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		s.logger.Warnf("can't access word meaning page: %s", err.Error())
		return nil, fmt.Errorf("Não foi possível acessar o significado da palavra desejada no Dicio.")
	}

	return meanings, nil
}

func (s *ScrapService) Synonyms(word string) ([]string, error) {
	c := s.scrap.GetColl()
	syns := []string{}
	c.OnHTML(".sinonimos", func(e *colly.HTMLElement) {
		if !strings.Contains(e.Text, "sinônimo") {
			return
		}
		synsSepByComma := strings.Split(e.Text, ":")
		if len(synsSepByComma) < 2 {
			return
		}
		for _, syn := range strings.Split(synsSepByComma[1], ", ") {
			syn := strings.Replace(syn, "\\n", "", -1)
			syn = strings.TrimSpace(syn)
			syns = append(syns, syn)
		}
	})

	err := c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		s.logger.Warnf("can't open dicio page of word %s: %s", word, err.Error())
		return nil, fmt.Errorf("Não foi possível encontrar a palavra no Dicio.")
	}
	return syns, nil
}
