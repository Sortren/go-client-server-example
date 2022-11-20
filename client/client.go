package main

import (
	"github.com/Sortren/go-client-server-example/stream"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("can't dial on localhost:8080, err: %s", err)
	}

	if err := stream.Write(conn, "[client] hello server!"); err != nil {
		log.Fatalf("can't write to the stream, err: %s", err)
	}

	if err := stream.Read(conn); err != nil {
		log.Fatalf("can't read from the stream, err: %s", err)
	}

}
