package powerdns

import (
	"github.com/cynt4k/router-dns-bridge/cmd/config"
	"github.com/cynt4k/router-dns-bridge/pkg/dns"
)

type PowerDNS interface {
	CreateRecordV4(record *dns.ARecord, router config.Router) error
	CreateRecordV6(record *dns.AAARecord, router config.Router) error
}
