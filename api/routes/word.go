package routes

import (
	"github.com/bruceneco/dicio-api/api/controllers"
	"github.com/bruceneco/dicio-api/lib"
)

type WordRoutes struct {
	logger         *lib.Logger
	handler        *lib.RequestHandler
	wordController *controllers.WordController
}

func (pr *WordRoutes) Setup() {
	pr.logger.Info("Setting up routes")

	api := pr.handler.Gin.Group("/word")
	{
		api.GET("/top-words", pr.wordController.GetTopWords)
		api.GET("/meanings/:word", pr.wordController.GetMeanings)
		api.GET("/synonyms/:word", pr.wordController.GetSynonyms)
	}
}

func NewWordRoutes(
	logger *lib.Logger,
	handler *lib.RequestHandler,
	wordController *controllers.WordController,
) *WordRoutes {
	return &WordRoutes{
		logger:         logger,
		handler:        handler,
		wordController: wordController,
	}
}
