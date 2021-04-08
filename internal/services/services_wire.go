// +build wireinject

package services

import (
	"github.com/google/wire"
)

// ProviderSet : Services dependency injection
var ProviderSet = wire.NewSet(wire.FieldsOf(new(*Services),
	"PowerDNS",
))
