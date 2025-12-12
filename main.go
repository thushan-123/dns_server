package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"

	"dns/encodeDecode"
	"dns/header"
)

func parseDomain(msg []byte, offset int) (string, int) {
	var domain string
	for {
		l := int(msg[offset])
		if l == 0 {
			offset++
			break
		}
		offset++
		domain += string(msg[offset:offset+l]) + "."
		offset += l
	}
	return domain, offset
}

func buildARecord(queryID uint16, domain string, ip string) []byte {
	// Header
	respHeader := header.DNSHeader{
		ID:      queryID,
		Flags:   0x8180, // standard response, no error
		QDCount: 1,
		ANCount: 1,
		NSCount: 0,
		ARCount: 0,
	}

	msg := header.CreateHeader(respHeader)

	// Question section (repeat same query)
	msg = append(msg, encodeDecode.EncodeDomain(domain)...)
	qtype := make([]byte, 4)
	binary.BigEndian.PutUint16(qtype[0:2], 1) // type A
	binary.BigEndian.PutUint16(qtype[2:4], 1) // class IN
	msg = append(msg, qtype...)

	// Answer section
	msg = append(msg, encodeDecode.EncodeDomain(domain)...)

	atype := make([]byte, 10)
	binary.BigEndian.PutUint16(atype[0:2], 1)     // type A
	binary.BigEndian.PutUint16(atype[2:4], 1)     // class IN
	binary.BigEndian.PutUint32(atype[4:8], 60)    // TTL
	binary.BigEndian.PutUint16(atype[8:10], 4)    // IPv4 length
	msg = append(msg, atype...)

	ipBytes := net.ParseIP(ip).To4()
	msg = append(msg, ipBytes...)

	return msg
}

func main() {
	buffer := make([]byte, 4096)

	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("DNS Server running on 127.0.0.1:1234")

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}

		msg := buffer[:n]

		//header
		id := binary.BigEndian.Uint16(msg[0:2])
		offset := 12

		//domain
		domain, _ := parseDomain(msg, offset)
		domain = domain[:len(domain)-1] 

		fmt.Println("Query for:", domain)

		//response
		resp := buildARecord(id, domain, "127.0.0.1")

		conn.WriteToUDP(resp, clientAddr)
	}
}


// dig @127.0.0.1 -p 1234 test.com A
