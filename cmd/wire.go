// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"github.com/jackkenney/gl-server/api"
)

func InjectServer() (*api.Server, error) {
	wire.Build(api.DefaultProviderSet)
	return &api.Server{}, nil
}
