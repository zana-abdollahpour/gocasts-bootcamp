package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	const (
		network = "tcp"
		address = "127.0.0.1:"
		port    = "8000"
	)

	message := "default response, since nothing is provided"
	if len(os.Args) >= 2 {
		message = os.Args[1]
	}

	connection, dialError := net.Dial(network, address+port)

	if dialError != nil {
		log.Fatalln("can't dial the given address", dialError)
	}

	numWrittenBytes, writeErr := connection.Write([]byte(message))

	if writeErr != nil {
		log.Fatalln("can't write data to connection", writeErr)
	}

	fmt.Println("numWrittenBytes", numWrittenBytes)

	var data = make([]byte, 1024)
	_, readError := connection.Read(data)

	if readError != nil {
		log.Fatal("can't read data from connection")
	}

	fmt.Printf("serverResponse: %s\n", string(data))
}
