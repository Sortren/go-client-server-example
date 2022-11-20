package main

import (
	"fmt"
	"github.com/Sortren/go-client-server-example/stream"
	"log"
	"net"
)

func main() {
	srv, err := net.Listen("tcp", ":8080")
	fmt.Println("listening on localhost:8080")

	if err != nil {
		log.Fatalf("can't listen on localhost:8080 err: %s", err)
	}

	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Printf("can't accept connection from: %s, err: %s", conn, err)
		}

		if err := stream.Read(conn); err != nil {
			log.Fatalf("can't read from stream, err: %s", err)
		}

		if err := stream.Write(conn, "[server] hello client!\n"); err != nil {
			log.Fatalf("can't write to the stream, err: %s", err)
		}
	}

}
