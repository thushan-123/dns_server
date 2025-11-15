package config

type DNSRecord struct {
	Name  string
	Type  uint16
	Class uint16
	TTL   uint32
	Data  string
}
