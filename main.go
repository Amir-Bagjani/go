package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "files/data.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([]byte, 20)

	_, rErr := file.Read(data)
	if rErr != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	fmt.Println(string(data))

}
