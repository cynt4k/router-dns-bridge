// +build wireinject

package router

import (
	"github.com/cynt4k/router-dns-bridge/cmd/config"
	v1 "github.com/cynt4k/router-dns-bridge/internal/router/v1"
	"github.com/cynt4k/router-dns-bridge/internal/services"
	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	"github.com/google/wire"
	"github.com/leandro-lugaresi/hub"
)

func newRouter(hub *hub.Hub, ss *services.Services, config *config.Config, logger logger.Logger) *router {
	wire.Build(
		newEcho,
		services.ProviderSet,
		wire.Struct(new(v1.Handlers), "*"),
		wire.Struct(new(router), "*"),
	)
	return nil
}
