package api

import (
	"github.com/google/wire"
	"github.com/jackkenney/gl-server/api/handler"
	"github.com/jackkenney/gl-server/model"
)

// DefaultProviderSet injects the dependencies from the Model through to the Router.
var DefaultProviderSet = wire.NewSet(
	wire.Bind(new(model.Model), new(model.MockModel)),
	model.ProvideMockModel,
	handler.ProvideModelHandler,
	ProvideRouter,
	ProvideServer,
)
