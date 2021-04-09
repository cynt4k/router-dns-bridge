package powerdns

import (
	"fmt"

	"github.com/cynt4k/router-dns-bridge/cmd/config"
	"github.com/cynt4k/router-dns-bridge/pkg/dns"
	pdnssdk "github.com/joeig/go-powerdns/v2"
)

func (s *service) createRecordPrecheck(router *config.Router) error {
	if !s.checkEnabled() {
		return nil
	}
	if router.Provider != "powerdns" {
		return fmt.Errorf("router provider does not match with powerdns: %s", router.Provider)
	}
	return nil
}

func (s *service) CreateRecordV4(record *dns.ARecord, router config.Router) error {
	if err := s.createRecordPrecheck(&router); err != nil {
		return err
	}

	if err := s.pdns.Records.Change(
		router.Domain,
		record.Name,
		pdnssdk.RRTypeA,
		record.TTL,
		[]string{record.IP.String()},
	); err != nil {
		return err
	}
	return nil
}

func (s *service) CreateRecordV6(record *dns.AAARecord, router config.Router) error {
	if err := s.createRecordPrecheck(&router); err != nil {
		return err
	}

	if !record.ValidIPv6() {
		return fmt.Errorf("no valid ipv6 address - could not be added to powerdns")
	}

	if err := s.pdns.Records.Change(
		router.Domain,
		record.Name,
		pdnssdk.RRTypeAAAA,
		record.TTL,
		[]string{record.IP.String()},
	); err != nil {
		return err
	}
	return nil
}
