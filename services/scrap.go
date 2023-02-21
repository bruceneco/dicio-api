package services

import (
	"github.com/bruceneco/dicio-api/lib"
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
	return []string{"Gato", "Cachorro"}, nil
}
