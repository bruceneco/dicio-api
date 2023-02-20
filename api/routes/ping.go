package routes

import (
	"github.com/bruceneco/dicio-api/api/controllers"
	"github.com/bruceneco/dicio-api/lib"
)

type PingRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	pingController controllers.PingController
}

func (pr *PingRoutes) Setup() {
	pr.logger.Info("Setting up routes")

	api := pr.handler.Gin.Group("/ping")
	{
		api.GET("", pr.pingController.Ping)
	}
}

func NewPingRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	pingController controllers.PingController,
) *PingRoutes {
	return &PingRoutes{
		logger:         logger,
		handler:        handler,
		pingController: pingController,
	}
}
