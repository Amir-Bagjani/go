package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer listen.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, cErr := listen.Accept()
		if cErr != nil {
			fmt.Println("cError:", err)
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 1024)

	// for {
	n, e := conn.Read(data)
	if e != nil {
		fmt.Println("Read Error:", e)

		return
	}

	// fmt.Println(string(data))
	fmt.Printf("Received: %s\n", data[:n])

	_, e = conn.Write([]byte("msg received"))
	if e != nil {
		fmt.Println("Error:", e)
	}
	// }
}
