package main

import (
    "bufio"
    "os"
)

func main() {
    f := os.Stdout

    w := bufio.NewWriterSize(f, 3)

    // Will print abc. d is added the buffer and will not be printed without an explicit .Flush()
    w.WriteString("abc")
    w.WriteString("b")
    // w.WriteString("c")
    // -----------------
    // w.WriteString("d")
	w.Flush()
}