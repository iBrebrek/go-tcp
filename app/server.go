package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	fmt.Println("Listening")

	for {
		conn, err := l.Accept()
		fmt.Println("Accepted")
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		handleConnection(conn)
		fmt.Println("Handled")
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	request := string(buffer)
	fmt.Println("~~~ request ~~~\n", request)
	requestParts := strings.Split(request, "\r\n")
	requestLine := requestParts[0]
	path := strings.Split(requestLine, " ")[1]
	if path == "/" || path == "/index.html" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nyes"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\nno"))
	}
}
