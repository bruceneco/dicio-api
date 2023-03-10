package controllers

import (
	"fmt"
	"github.com/bruceneco/dicio-api/services"
	"github.com/bruceneco/dicio-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WordController struct {
	scrap *services.ScrapService
}

func NewWordController(scrap *services.ScrapService) *WordController {
	return &WordController{scrap: scrap}
}

func (wc *WordController) GetTopWords(c *gin.Context) {
	nWords := 200
	queryNWords, exist := c.GetQuery("nWords")
	if exist {
		var err error
		nWords, err = strconv.Atoi(queryNWords)
		if err != nil {
			utils.NewError(c, http.StatusBadRequest, "Não foi possível ler a quantidade de palavras desejada.")
			return
		}
	}
	words, err := wc.scrap.TopWords(nWords)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"topWords":   words,
		"wordsCount": len(words),
	})
}

func (wc *WordController) GetMeanings(c *gin.Context) {
	word := c.Param("word")
	meanings, err := wc.scrap.Meanings(word)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"meanings": meanings})
}

func (wc *WordController) GetSynonyms(c *gin.Context) {
	word := c.Param("word")
	syns, err := wc.scrap.Synonyms(word)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	} else if len(syns) == 0 {
		utils.NewError(c, http.StatusNotFound, fmt.Sprintf("Não há sinônimos para a palavra \"%s\".", word))
		return
	}
	c.JSON(http.StatusOK, gin.H{"synonyms": syns})
}

func (wc *WordController) GetAntonyms(c *gin.Context) {
	word := c.Param("word")
	antonyms, err := wc.scrap.Antonyms(word)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	} else if len(antonyms) == 0 {
		utils.NewError(c, http.StatusNotFound, fmt.Sprintf("Não há antônimos para a palavra \"%s\".", word))
		return
	}
	c.JSON(http.StatusOK, gin.H{"antonyms": antonyms})
}

func (wc *WordController) GetEtymology(c *gin.Context) {
	word := c.Param("word")
	etym, err := wc.scrap.Etymology(word)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	} else if etym == "" {
		utils.NewError(c, http.StatusNotFound, fmt.Sprintf("Não há sinônimos para a palavra \"%s\".", word))
		return
	}
	c.JSON(http.StatusOK, gin.H{"etymology": etym})
}

func (wc *WordController) GetDefinition(c *gin.Context) {
	word := c.Param("word")
	def, err := wc.scrap.Definition(word)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	} else if def == nil {
		utils.NewError(c, http.StatusNotFound, fmt.Sprintf("Não foi possível encontrar a definição de %s.", word))
		return
	}
	c.JSON(http.StatusOK, def)
}

func (wc *WordController) GetExamples(c *gin.Context) {
	word := c.Param("word")
	exs, err := wc.scrap.Examples(word)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	} else if len(exs) == 0 {
		utils.NewError(c, http.StatusNotFound, fmt.Sprintf("Não foi possível encontrar exemplos utilizando a palavra \"%s\".", word))
		return
	}
	c.JSON(http.StatusOK, gin.H{"examples": exs})
}

func (wc *WordController) GetCitations(c *gin.Context) {
	word := c.Param("word")
	citations, err := wc.scrap.Citations(word)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	} else if len(citations) == 0 {
		utils.NewError(c, http.StatusNotFound, fmt.Sprintf("Não foi possível encontrar citações utilizando a palavra \"%s\".", word))
		return
	}
	c.JSON(http.StatusOK, gin.H{"citations": citations})
}

func (wc *WordController) GetFullInfo(c *gin.Context) {
	word := c.Param("word")
	wi, err := wc.scrap.FullWordInfo(word)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	} else if wi == nil {
		utils.NewError(c, http.StatusNotFound, fmt.Sprintf("Não foi possível encontrar a palavra \"%s\".", word))
		return
	}
	c.JSON(http.StatusOK, wi)
}
