package stream

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
)

func ReadExperimental(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(conn)
	var str strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		str.WriteString(line)
	}

	return str.String(), nil
}

func Read(conn net.Conn) error {
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer

	for {
		line, isPrefix, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		buffer.Write(line)

		if !isPrefix {
			break
		}
	}

	fmt.Println(buffer.String())

	return nil
}

func Write(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))

	if err != nil {
		return fmt.Errorf("err while Writing into conn: %w", err)
	}

	fmt.Println(message)

	return nil
}
