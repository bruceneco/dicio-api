package commands

import (
	"github.com/bruceneco/dicio-api/api/middlewares"
	"github.com/bruceneco/dicio-api/api/routes"
	"github.com/bruceneco/dicio-api/lib"
	"github.com/spf13/cobra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(_ *cobra.Command) {}

func (s *ServeCommand) Run() lib.CommandRunner {
	return func(
		middleware *middlewares.Middlewares,
		env *lib.Env,
		router *lib.RequestHandler,
		route *routes.Routes,
		logger *lib.Logger,
	) {
		middleware.Setup()
		route.Setup()

		logger.Info("Running server")
		if env.ServerPort == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.ServerPort)
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
