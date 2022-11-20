package stream

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
)

func Read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer

	for {
		line, isPrefix, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		buffer.Write(line)

		if !isPrefix {
			break
		}
	}
	return buffer.String(), nil
}

func Write(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))

	if err != nil {
		return fmt.Errorf("err while Writing into conn: %w", err)
	}

	return nil
}
