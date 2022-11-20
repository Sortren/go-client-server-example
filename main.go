package main

import (
	"bufio"
	"net"
	"strings"
)

func TCPStreamScannerBufio(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(conn)
	var str strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		str.WriteString(line)
	}

	return str.String(), nil
}

func main() {

}
