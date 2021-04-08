package powerdns

import (
	"fmt"

	"github.com/cynt4k/router-dns-bridge/cmd/config"
	"github.com/cynt4k/router-dns-bridge/pkg/logger"
	pdnssdk "github.com/joeig/go-powerdns/v2"
	"github.com/leandro-lugaresi/hub"
)

var (
	s *service
)

type service struct {
	hub  *hub.Hub
	pdns *pdnssdk.Client
}

func NewService(hub *hub.Hub, config *config.Config, logger logger.Logger) (PowerDNS, error) {
	if s != nil {
		return s, nil
	}
	headers := map[string]string{"X-API-Key": config.Providers.Powerdns.APIKey}
	pdns := pdnssdk.NewClient(config.Providers.Powerdns.URL, "localhost", headers, nil)
	s = &service{
		hub:  hub,
		pdns: pdns,
	}
	_, err := pdns.Servers.List()
	if err != nil {
		return nil, fmt.Errorf("could not connect to PowerDNS servers: %w", err)
	}
	return nil, nil
}
