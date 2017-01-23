package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", os.Args[1])

	if err != nil {
		log.Fatal("Error resolving ip address and/or port: ", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)

	if err != nil {
		log.Fatal("Error connecting to remote address: ", err)
	}

	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")
		packet, err := reader.ReadBytes('\n')

		if err != nil {
			log.Fatal("Error reading input: ", err)
		}

		bytes, err := conn.Write(packet)

		if err != nil {
			log.Fatal("Error writing input to UDP packet: ", err)
		}

		log.Printf("%d bytes of data sent to server", bytes)
	}
}
