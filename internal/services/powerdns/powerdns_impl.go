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
	enabled bool
	hub     *hub.Hub
	pdns    *pdnssdk.Client
	config  *config.Config
	logger  logger.Logger
}

func NewService(hub *hub.Hub, config *config.Config, logger logger.Logger) (PowerDNS, error) {
	log := logger.New("powerdns")
	if s != nil {
		return s, nil
	}

	if config.Providers.Powerdns == nil {
		log.Info("no powerdns provider - disabled")
		return &service{
			enabled: false,
			hub:     hub,
			config:  config,
			logger:  log,
		}, nil
	}

	headers := map[string]string{"X-API-Key": config.Providers.Powerdns.APIKey}
	pdns := pdnssdk.NewClient(config.Providers.Powerdns.URL, "localhost", headers, nil)
	s = &service{
		enabled: true,
		hub:     hub,
		pdns:    pdns,
		config:  config,
		logger:  log,
	}
	_, err := pdns.Servers.List()
	if err != nil {
		return nil, fmt.Errorf("could not connect to PowerDNS servers: %w", err)
	}
	return s, nil
}

func (s *service) checkEnabled() bool {
	if !s.enabled {
		s.logger.Warn("no powerdns action will be executed because its not enabled")
	}
	return s.enabled
}
