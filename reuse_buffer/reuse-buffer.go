package reuse_buffer

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const singleLine string = "I'd love to have some coffee right about now"

const multiLine string = "Reading is my...\r\n favourite"

func ReuseBuffer() {

	fmt.Println("Length of singleLine input is " + strconv.Itoa(len(singleLine)))

	str := strings.NewReader(singleLine)
	bf := bufio.NewReaderSize(str, 25)

	fmt.Println("\n---Peek---")
	// Peek - Case 1: Simple peek implementation
	b, err := bf.Peek(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("3 char ==> %q \n\n", b)

	// Peek - Case 2: Peek larger than buffer size
	b, err = bf.Peek(300)
	if err != nil {
		fmt.Println(err) // output: "bufio: buffer full"
	}
	fmt.Printf("300 char ==> %q \n\n", b) // b size is equal to NewReaderSize, in this case 25

	// Peek - Case3: buffer size larger than string (buffer is not full)
	bf_large := bufio.NewReaderSize(str, 500)
	b, err = bf_large.Peek(500)
	if err != nil {
		fmt.Println(err) // output: EOF
	}
	fmt.Printf("%q \n\n", b)

	// ReadSlice
	fmt.Println("\n---ReadSlice---")
	str = strings.NewReader(multiLine)
	bs := bufio.NewReader(str)
	for {
		token, err := bs.ReadSlice('.')
		if len(token) > 0 {
			fmt.Printf("Token (ReadSlice): %q\n", token)
		}
		if err != nil {
			break
		}
	}

	// ReadLine
	fmt.Println("\n---ReadLine---")
	str = strings.NewReader(multiLine)
	bs = bufio.NewReader(str)
	for {
		token, _, err := bs.ReadLine()
		if len(token) > 0 {
			fmt.Printf("Token (ReadLine): %q\n", token)
		}
		if err != nil {
			break
		}
	}

	// ReadBytes
	fmt.Println("\n---ReadBytes---")
	str = strings.NewReader(multiLine)
	bs.Reset(str)
	for {
		token, err := bs.ReadBytes('\n')
		fmt.Printf("Token (ReadBytes): %q\n", token)
		if err != nil {
			break
		}
	}

	// Scanner
	fmt.Println("\n---Scanner---")
	str = strings.NewReader(multiLine)
	scanner := bufio.NewScanner(str)

	for scanner.Scan() {
		fmt.Printf("Token (Scanner): %q\n", scanner.Text())
	}

}
