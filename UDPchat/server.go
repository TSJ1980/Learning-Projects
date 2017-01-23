package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal("Error converting string port to integer: ", err)
	}

	addr := net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: port,
	}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		log.Fatal("Error listening on port: ", err)
	}

	log.Printf("Listening on port %d", port)

	for true {
		byte_message := make([]byte, 65, 507)
		bytes, returnAddr, err := conn.ReadFromUDP(byte_message)

		if err != nil {
			log.Fatal("Error reading UDP: ", err)
		}

		message := string(byte_message)
		log.Printf("%d bytes sent from address %s", bytes, returnAddr)
		log.Printf("%s", message)
	}
}
