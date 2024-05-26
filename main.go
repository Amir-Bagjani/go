package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.CreateTemp("", "sample_")
	check(err)
	defer os.Remove(file.Name())

	fmt.Println("temporary file name is", file.Name())

	_, err = file.Write([]byte("hello mo"))
	check(err)

	dir, dErr := os.MkdirTemp("", "sample_dir_")
	check(dErr)
	defer os.RemoveAll(dir)

	fmt.Println("Temp dir name:", dir)

	p := filepath.Join(dir, "data.txt")

	err = os.WriteFile(p,[]byte("asdasdasd"), 0644)
	check(err)
}
