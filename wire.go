//+build wireinject

package main

import "github.com/google/wire"

func InitializeApp() (*App, func(), error) {
	wire.Build(
		NewListener,
		NewServer,
		App{},
	)
	return &App{}, nil, nil
}
