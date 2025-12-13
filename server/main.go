package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		network = "tcp"
		address = "127.0.0.1:"
		port    = "8000"
	)

	// Create new listener
	listener, listenErr := net.Listen(network, address+port)

	if listenErr != nil {
		log.Fatal("Can't listen to the given address", address, listenErr)
	}

	defer listener.Close()

	givenAddress := listener.Addr()

	fmt.Println(givenAddress)

	for {
		// listen for new connection
		connection, connectionErr := listener.Accept()

		if connectionErr != nil {
			log.Println("Can't listen to new connection", connectionErr)
		}

		// process request
		var data = make([]byte, 1024)
		numReadBytes, readError := connection.Read(data)

		if readError != nil {
			log.Println("can't read data from connection")

			continue
		}

		fmt.Printf("numReadBytes %d\ndata: %+v\n", numReadBytes, string(data))

		_, writeErr := connection.Write([]byte(`Your message has been received, your Majesty!`))

		if writeErr != nil {
			log.Println("can't write data to connection")

			continue
		}
	}

}
