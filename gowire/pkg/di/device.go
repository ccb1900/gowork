//go:build wireinject
// +build wireinject

package di

import (
	"gowire/pkg/handler"
	// "gowire/pkg/repo"

	"github.com/google/wire"
)

func InitializeDeviceHandler() handler.BookHandler {
	wire.Build(
		handler.NewBookHandler,
		bookSet,
	)
	return handler.BookHandler{}
}

// var bookSet = wire.NewSet(
// 	repo.NewBookRepo,
// 	wire.Bind((*repo.IBookRepo)(nil), (*repo.BookRepo)(nil)),
// )
