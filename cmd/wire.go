// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"time"

	"github.com/google/wire"
	"github.com/jackkenney/gl-server/api"
)

func InjectServer(wait time.Duration) (api.Server, error) {
	wire.Build(
		api.DefaultProviderSet,
		wire.Value(wait),
	)
	return api.Server{}, nil
}
