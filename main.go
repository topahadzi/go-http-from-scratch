package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	// Convert the request data to a string
	request := string(buffer)

	// Check for the "efishery-domain" header
	if strings.Contains(request, "efishery-domain: infra") {
		// Check if the request is for the root path ("/")
		if strings.Contains(request, "GET / ") || strings.Contains(request, "GET / HTTP/1.1") {
			// Respond with a 200 OK status and serve the "index.html" file
			response := "HTTP/1.1 200 OK\r\nContent-Length: 5\r\n\r\nEFisheryWeb"
			conn.Write([]byte(response))
		} else {
			// Respond with a 404 Not Found status for other paths
			response := "HTTP/1.1 404 Not Found\r\nContent-Length: 0\r\n\r\n"
			conn.Write([]byte(response))
		}
	} else {
		// Respond with a 503 Service Unavailable status
		response := "HTTP/1.1 503 Service Unavailable\r\nContent-Length: 0\r\n\r\n"
		conn.Write([]byte(response))
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listen.Close()
	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
