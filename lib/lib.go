package lib

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewEnv),
	fx.Provide(GetLogger),
	fx.Provide(NewScrap),
	fx.Provide(NewTextTransform),
)
