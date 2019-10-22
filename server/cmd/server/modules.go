//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/Tornado9966/menu-example/server/menu"
)

// ComposeApiServer will create an instance of MenuApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*MenuApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from dishes package.
		dishes.Providers,
		// Provide MenuApiServer instantiating the structure and injecting dishes handler and port number.
		wire.Struct(new(MenuApiServer), "Port", "DishesHandler"),
	)
	return nil, nil
}
