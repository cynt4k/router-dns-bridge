// +build wireinject

package cmd

import (
	"github.com/cynt4k/router-dns-bridge/cmd/config"
	"github.com/cynt4k/router-dns-bridge/internal/router"
	"github.com/cynt4k/router-dns-bridge/internal/services"
	"github.com/cynt4k/router-dns-bridge/internal/services/powerdns"
	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	"github.com/google/wire"
	"github.com/leandro-lugaresi/hub"
)

func newServer(hub *hub.Hub, logger logger.Logger, config *config.Config) (*Server, error) {
	wire.Build(
		powerdns.NewService,
		router.New,
		wire.Struct(new(services.Services), "*"),
		wire.Struct(new(Server), "*"),
	)
	return nil, nil
}
