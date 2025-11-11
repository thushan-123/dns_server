package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// memory buffer
	p := make([]byte, 4048)

	// create the udp socket listen port 53

	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1"),
	}

	connection, err := net.ListenUDP("udp", &addr)

	if err != nil {
		fmt.Print("error ", err)
		os.Exit(1)
	}

	defer connection.Close()

	for {
		_, clientAdress, err := connection.ReadFromUDP(p)

		if err != nil {
			fmt.Print("\nerror :", err)

		}
		s.WriteToUDP([]byte("hii"), clientAdress)
	}

}
