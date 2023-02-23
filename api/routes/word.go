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
		api.GET("/etymology/:word", pr.wordController.GetEtymology)
		api.GET("/definition/:word", pr.wordController.GetDefinition)
		api.GET("/examples/:word", pr.wordController.GetExamples)
		api.GET("/citations/:word", pr.wordController.GetCitations)
		api.GET("/antonyms/:word", pr.wordController.GetAntonyms)
		api.GET("/info/:word", pr.wordController.GetFullInfo)
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
