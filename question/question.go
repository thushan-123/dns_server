package question

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
