package question

import "encoding/binary"


type DnsHeader struct {
	ID      uint16 // alls 16 bits
	Falgs   uint16 //  QR- 1 | OPCODE- 4 | AA-1 |  TC -1 |  RD- 1 |  RA -1 |  Z- 3  | RCODE- 4
	RCODE   uint16
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

type DNSMessage []byte

func CreateHeader(dnsHeader DnsHeader) DNSMessage {
	header := make([]byte, 12)

	binary.BigEndian.PutUint16(header[0:2], dnsHeader.ID)
	binary.BigEndian.PutUint16(header[2:4], dnsHeader.Falgs)
	binary.BigEndian.PutUint16(header[4:6], dnsHeader.QDCOUNT)
	binary.BigEndian.PutUint16(header[6:8], dnsHeader.ANCOUNT)
	binary.BigEndian.PutUint16(header[8:10], dnsHeader.NSCOUNT)
	binary.BigEndian.PutUint16(header[10:12], dnsHeader.ARCOUNT)

	return header
}



/*
	ID - 16 bits
	QR - 1 bit
	OPCODE - 4 bit
	AA - 1 bit
	TC - 1 bit
	RD - 1 bit					16		16         16        16       16        16				16
	RA - 1 bit				|   ID   |         |   QDCOUNT  |   ANCOUNT     |   NSCOUNT     |  ARCOUNT     |
	Z  - 3 bit						 /          \
	RCODE - 4 bit			QR| OPCODE | AA | TC | RD | RA | Z|RCODE
	QDCOUNT - 16 bit
	ANCOUNT - 16 bit
	NSCOUNT - 16 bit
	ARCOUNT - 16 bit

*/
