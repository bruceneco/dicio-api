package controllers

import (
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
	c.JSON(http.StatusOK, words)
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
