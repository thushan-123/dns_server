package header

import "encoding/binary"

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

type DNSHeader struct {
	ID      uint16
	Flags   uint16
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

type DNSMessage []byte

func CreateHeader(h DNSHeader) DNSMessage {
	header := make([]byte, 12)

	binary.BigEndian.PutUint16(header[0:2], h.ID)
	binary.BigEndian.PutUint16(header[2:4], h.Flags)
	binary.BigEndian.PutUint16(header[4:6], h.QDCount)
	binary.BigEndian.PutUint16(header[6:8], h.ANCount)
	binary.BigEndian.PutUint16(header[8:10], h.NSCount)
	binary.BigEndian.PutUint16(header[10:12], h.ARCount)

	return header
}





