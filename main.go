package main

import (
	"fmt"
	"net"
	"os"
)

func SendMessage(connection net.Conn, clientAddress net.UDPAddr) {

}

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

	fmt.Println("Server Is Running on localhost port 1234")

	for {
		i, clientAdress, err := connection.ReadFromUDP(p)

		if err != nil {
			fmt.Print("\nerror :", err)

		}

		fmt.Printf("recived bytes %d  %s", i, clientAdress)

		fmt.Printf("data: %s\n", string(p[:i])) // print  revived data

		connection.WriteToUDP([]byte("hii"), clientAdress) // client respons befor send
	}

}
