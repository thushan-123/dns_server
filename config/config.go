package config

import "net"

// dns record type

const (
	TypeA     = 1
	TypeNS    = 2
	TypeCNAME = 5
	TypeSOA   = 6
	TypeMX    = 15
	TypeTXT   = 16
	TypeAAA   = 28
)

type DNSRecord struct {
	Name  string
	Type  uint16
	Class uint16
	TTL   uint32
	Data  string
}

// dns class
const (
	ClassIN = 1
)

type DNSConfig struct {
	Port int
	Forwarders [] string
	CacheTTL int
	RecordsFile string
}

type DNSServer struct{
	Config DNSConfig
	Records DNSRecord
	Listner net.PacketConn

}
