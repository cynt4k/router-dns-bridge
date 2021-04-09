package dns

import "net"

type ARecord struct {
	Name string
	IP   net.IP
	TTL  uint32
}

type AAARecord struct {
	Name string
	IP   net.IP
	TTL  uint32
}

type CNAMERecord struct {
	Name        string
	Destination string
	TTL         uint32
}

type TXTRecord struct {
	Name string
	Data string
	TTL  uint32
}

func (r *AAARecord) ValidIPv6() bool {
	return net.ParseIP(r.IP.String()).To16() != nil
}
