package bootstrap

import (
	"github.com/bruceneco/dicio-api/api/controllers"
	"github.com/bruceneco/dicio-api/api/middlewares"
	"github.com/bruceneco/dicio-api/api/routes"
	"github.com/bruceneco/dicio-api/lib"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	routes.Module,
	lib.Module,
	middlewares.Module,
	controllers.Module,
)
