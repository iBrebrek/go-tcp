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
	resp := dispatch(parseRequest(request))
	conn.Write([]byte(resp))
}

func parseRequest(request string) (method, path, body string) {
	fmt.Println("~~~ request ~~~\n", request)
	requestParts := strings.Split(request, "\r\n")
	requestLine := requestParts[0]
	lineParts := strings.Split(requestLine, " ")
	method = lineParts[0]
	path = lineParts[1]
	body = requestParts[2]
	return // named result parameters
}

func dispatch(method string, path string, body string) string {
	if method != "GET" {
		return "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
	}

	if path == "/" || path == "/index.html" {
		return "HTTP/1.1 200 OK\r\n\r\n"
	} else {
		return "HTTP/1.1 404 Not Found\r\n\r\n"
	}
}
