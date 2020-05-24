// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/jackkenney/gl-server/api"
	"github.com/jackkenney/gl-server/api/handler"
	"github.com/jackkenney/gl-server/model"
)

// Injectors from wire.go:

func InjectServer() (*api.Server, error) {
	mockModel := model.ProvideMockModel()
	modelHandler := handler.ProvideModelHandler(mockModel)
	router := api.ProvideRouter(modelHandler)
	server := api.ProvideServer(router)
	return server, nil
}
