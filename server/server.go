package main

import (
	"log"
	"net"
)

func main() {
	srv, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("can't listen on localhost:8080 err: %s", err)
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Printf("can't accept connection from: %s, err: %s", conn, err)
		}

		stream.Read()

	}

}
