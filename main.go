package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println(filepath.Base(p))
	fmt.Println(filepath.Dir(p))

	d, f := filepath.Split(p)
	fmt.Println("directory =>", d, "base is =>", f)

	fmt.Println(filepath.IsAbs("/dir"))
	fmt.Println(filepath.IsAbs("dir"))

	filename := "config.json"
	exc := filepath.Ext(filename)
	fmt.Println(exc)
	fmt.Println(strings.TrimSuffix(filename, exc))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
        panic(err)
    }
    fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
        panic(err)
    }
    fmt.Println(rel)

}
