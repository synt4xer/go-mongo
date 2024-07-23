//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/synt4xer/go-mongo/config"
	"github.com/synt4xer/go-mongo/internal/delivery/http"
	"github.com/synt4xer/go-mongo/internal/repository"
	"github.com/synt4xer/go-mongo/internal/usecase"
)

func apiHandlers() (*http.UserHandler, error) {
	panic(wire.Build(
		config.ProvideConfig,
		config.ProvideClient,
		repository.NewMongoRepository,
		repository.NewUserRepository,
		usecase.ProvideUserUseCase,
		http.NewUserHandler,
	))
}
