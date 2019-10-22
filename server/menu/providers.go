package menu

import "github.com/google/wire"

// Set of providers for dishes components.
var Providers = wire.NewSet(NewStore, HttpHandler)
