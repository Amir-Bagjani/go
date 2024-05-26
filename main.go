package main

import (
	"bufio"
	"fmt"
	"strings"
)

const str = "Reading is my...\r\n favourite"

func main() {
	fmt.Println("length of reader is", len(str))

	stringReader := strings.NewReader(str)

	bs := bufio.NewScanner(stringReader)

	for bs.Scan() {
		fmt.Printf("The output is %q \n", bs.Text())
	}
}
