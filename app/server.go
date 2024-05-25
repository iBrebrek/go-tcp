package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	fmt.Println("Listening")

	conn, err := l.Accept()
	fmt.Println("Accepted")
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	handleConnection(conn)
	fmt.Println("Handled")
}

func handleConnection(conn net.Conn) {
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
