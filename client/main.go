package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Send a sentence with 8 words to the server
	message := "This is a test sentence with eight words"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Close the write side of the connection to indicate we're done sending
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
	}

	data := make([]byte, 0, 4096)
	temp := make([]byte, 4096)

	for {
		fmt.Println("reading")
		n, err := conn.Read(temp)
		if err != nil {
			if err == io.EOF {
				fmt.Println("End of file.")
			} else {
				fmt.Println("Error reading from connection:", err)
			}
			break
		}

		data = append(data, temp[:n]...)
	}

	fmt.Println("Received from server:", string(data))
}
