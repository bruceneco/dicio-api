package services

import (
	"fmt"
	"github.com/bruceneco/dicio-api/lib"
	"github.com/bruceneco/dicio-api/models"
	"github.com/gocolly/colly"
	"math"
	"regexp"
	"strings"
	"time"
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
	s.extractMeaningsFromPage(c, &meanings)

	err = c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		s.logger.Warnf("can't access word meaning page: %s", err.Error())
		return nil, fmt.Errorf("Não foi possível acessar o significado da palavra desejada no Dicio.")
	}

	return meanings, nil
}

func (s *ScrapService) extractMeaningsFromPage(c *colly.Collector, meanings *[]*models.Meaning) {
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
		*meanings = append(*meanings, &meaning)
	})
}

func (s *ScrapService) Synonyms(word string) ([]string, error) {
	word, err := s.textTransform.RemoveAccents(strings.ToLower(word))
	if err != nil {
		s.logger.Warnf("can't remove accents from word: %s", err.Error())
		return nil, fmt.Errorf("Não foi possível remover os acentos da palavra para uma busca precisa.")
	}
	c := s.scrap.GetColl()
	syns := []string{}
	s.extractSynonymsFromPage(c, &syns)

	err = c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		s.logger.Warnf("can't open dicio page of word %s: %s", word, err.Error())
		return nil, fmt.Errorf("não foi possível encontrar sinônimos de %s.", word)
	}
	return syns, nil
}

func (s *ScrapService) extractSynonymsFromPage(c *colly.Collector, syns *[]string) {
	c.OnHTML(".sinonimos", func(e *colly.HTMLElement) {
		if !strings.Contains(e.Text, "sinônimo") {
			return
		}
		synsSepByComma := strings.Split(e.Text, ":")
		if len(synsSepByComma) < 2 {
			return
		}
		for _, syn := range strings.Split(synsSepByComma[1], ", ") {
			syn := strings.Replace(syn, "\n", "", -1)
			syn = strings.TrimSpace(syn)
			*syns = append(*syns, syn)
		}
	})
}

func (s *ScrapService) Etymology(word string) (string, error) {
	word, err := s.textTransform.RemoveAccents(strings.ToLower(word))
	if err != nil {
		s.logger.Warnf("can't remove accents from word: %s", err.Error())
		return "", fmt.Errorf("Não foi possível remover os acentos da palavra para uma busca precisa.")
	}
	c := s.scrap.GetColl()
	etym := ""
	s.extractEtymologyFromPage(c, &etym)

	err = c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		return "", fmt.Errorf("Não foi possível buscar a etimologia de %s.", word)
	}
	return etym, nil
}

func (s *ScrapService) extractEtymologyFromPage(c *colly.Collector, etym *string) {
	c.OnHTML(".significado > .etim", func(element *colly.HTMLElement) {
		*etym = strings.Split(element.Text, "). ")[1]
	})
}

func (s *ScrapService) Definition(word string) (*models.Definition, error) {
	word, err := s.textTransform.RemoveAccents(strings.ToLower(word))
	if err != nil {
		s.logger.Warnf("can't remove accents from word: %s", err.Error())
		return nil, fmt.Errorf("Não foi possível remover os acentos da palavra para uma busca precisa.")
	}
	c := s.scrap.GetColl()
	def := models.Definition{}
	c.OnHTML(".adicional", func(element *colly.HTMLElement) {
		txt := element.Text
		grammClass, err := s.textTransform.GetFromSubstrUntilTheEOL(txt, "Classe gramatical: ")
		if err == nil {
			def.GrammClass = grammClass
		}
		syllabicSep, err := s.textTransform.GetFromSubstrUntilTheEOL(txt, "Separação silábica: ")
		if err == nil {
			def.SyllabicSep = strings.Split(syllabicSep, "-")
		}
		plural, err := s.textTransform.GetFromSubstrUntilTheEOL(txt, "Plural: ")
		if err == nil {
			def.Plural = plural
		}
	})

	err = c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		return nil, fmt.Errorf("Não foi possível buscar a definição de %s.", word)
	}

	return &def, nil
}

func (s *ScrapService) Examples(word string) ([]*models.Example, error) {
	word, err := s.textTransform.RemoveAccents(strings.ToLower(word))
	if err != nil {
		s.logger.Warnf("can't remove accents from word: %s", err.Error())
		return nil, fmt.Errorf("Não foi possível remover os acentos da palavra para uma busca precisa.")
	}
	c := s.scrap.GetColl()
	examples := []*models.Example{}

	c.OnHTML(".frase", func(e *colly.HTMLElement) {
		re := regexp.MustCompile(`(\d{2}/\d{2}/\d{4})`)
		txt := e.Text
		if re.FindString(txt) == "" {
			return
		}
		example := models.Example{}
		authorAndDate := e.ChildText("em")
		parts := strings.Split(authorAndDate, ", ")
		example.Author = strings.Join(parts[:len(parts)-1], ", ")
		date, err := time.Parse("02/01/2006", parts[len(parts)-1])
		if err == nil {
			example.Date = date
		}

		example.Content = strings.Replace(txt, authorAndDate, "", -1)
		example.Content = strings.Replace(example.Content, "\n", "", -1)
		example.Content = strings.TrimSpace(example.Content)
		examples = append(examples, &example)
	})

	err = c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		s.logger.Warnf("could not find usage examples of \"%s\": %s", word, err.Error())
		return nil, fmt.Errorf("Não foi possível buscar exemplos de \"%s\".", word)
	}
	return examples, nil
}

func (s *ScrapService) Citations(word string) ([]*models.Citation, error) {
	word, err := s.textTransform.RemoveAccents(strings.ToLower(word))
	if err != nil {
		s.logger.Warnf("can't remove accents from word: %s", err.Error())
		return nil, fmt.Errorf("Não foi possível remover os acentos da palavra para uma busca precisa.")
	}
	c := s.scrap.GetColl()
	citations := []*models.Citation{}

	c.OnHTML(".frase", func(e *colly.HTMLElement) {
		re := regexp.MustCompile(`(\d{2}/\d{2}/\d{4})`)
		txt := e.Text
		if re.FindString(txt) != "" {
			return
		}
		citation := models.Citation{}
		rawAuthor := e.ChildText("em")
		citation.Author = strings.TrimSpace(strings.Replace(rawAuthor, "- ", "", -1))
		citation.Content = strings.Replace(txt, rawAuthor, "", -1)
		citation.Content = strings.Replace(citation.Content, "\n", "", -1)
		citation.Content = strings.TrimSpace(citation.Content)
		citations = append(citations, &citation)
	})

	err = c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		s.logger.Warnf("could not find usage citations of \"%s\": %s", word, err.Error())
		return nil, fmt.Errorf("Não foi possível buscar citações de \"%s\".", word)
	}
	return citations, nil
}

func (s *ScrapService) Antonyms(word string) ([]string, error) {
	word, err := s.textTransform.RemoveAccents(strings.ToLower(word))
	if err != nil {
		s.logger.Warnf("can't remove accents from word: %s", err.Error())
		return nil, fmt.Errorf("Não foi possível remover os acentos da palavra para uma busca precisa.")
	}
	c := s.scrap.GetColl()
	antonyms := []string{}
	c.OnHTML(".sinonimos", func(e *colly.HTMLElement) {
		if !strings.Contains(e.Text, "contrário") {
			return
		}
		antonymsSepByComma := strings.Split(e.Text, ":")
		if len(antonymsSepByComma) < 2 {
			return
		}
		for _, antonym := range strings.Split(antonymsSepByComma[1], ", ") {
			syn := strings.Replace(antonym, "\n", "", -1)
			syn = strings.TrimSpace(syn)
			antonyms = append(antonyms, syn)
		}
	})

	err = c.Visit(fmt.Sprintf("%s/%s", dicioURL, word))
	if err != nil {
		s.logger.Warnf("can't open dicio page of word %s: %s", word, err.Error())
		return nil, fmt.Errorf("não foi possível encontrar antônimos de %s.", word)
	}
	return antonyms, nil
}
