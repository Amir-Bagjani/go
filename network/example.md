### example of send data and received back the data

#### server 
``` go
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// Listen on TCP port 8080 on all available interfaces
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

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

	fmt.Println("data that received", string(data))

	// Write the received data back to the client
	_, err := conn.Write(data)
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}

	fmt.Println("Data sent back to client")
}
```

### client

``` go

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
```