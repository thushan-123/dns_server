package question

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

type Question struct {
	ID      byte
	QR      byte
	OPCODE  byte
	AA      byte
	TC      byte
	RD      byte
	RA      byte
	Z       byte
	RCODE   byte
	QDCOUNT byte
	ANCOUNT byte
	NSCOUNT byte
	ARCOUNT byte
}
