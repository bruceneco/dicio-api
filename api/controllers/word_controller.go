package controllers

import (
	"github.com/bruceneco/dicio-api/services"
	"github.com/bruceneco/dicio-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WordController struct {
	scrap *services.ScrapService
}

func NewWordController(scrap *services.ScrapService) *WordController {
	return &WordController{scrap: scrap}
}

func (wc *WordController) GetTopWords(c *gin.Context) {
	words, err := wc.scrap.TopWords()
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, words)
}
