package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func foo(w *io.PipeWriter) {
	defer w.Close()
	// Write a message to pipe writer
	fmt.Fprintln(w, "yooo ten")
}
func Io() {
	ten()
}

func ten() {
	pr, pw := io.Pipe()

	// Pass writer to function
	go foo(pw)

	var b bytes.Buffer
	mw := io.MultiWriter(&b, os.Stdout)

	_, err := io.Copy(mw, pr)
	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())
}

func nine() {
	r, w := io.Pipe()

	go func() {
		defer w.Close()
		fmt.Fprintln(w, "yoo nine")
	}()

	msg, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(msg))
}

func eight() {
	r := strings.NewReader("yooo eight")

	var b bytes.Buffer

	_, err := io.Copy(&b, r)
	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())
}

func seven() {
	r := strings.NewReader("yooo seven")

	var b bytes.Buffer

	b.ReadFrom(r)

	fmt.Println(b.String())
}

func one() {
	fmt.Fprintln(os.Stdout, "yoo std out")
}

func two() {
	var b bytes.Buffer

	fmt.Fprintln(&b, "yooo two")
	fmt.Println(b.String())
}

func three() {
	var a bytes.Buffer
	var b bytes.Buffer

	mb := io.MultiWriter(&a, &b)

	fmt.Fprintln(mb, "yooo three")
	fmt.Println(a.String())
	fmt.Println(b.String())
}

func four() {
	str := strings.NewReader("yooo four")

	msg, err := io.ReadAll(str)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(msg))
}

func five() {
	a := strings.NewReader("yooo five1")
	b := strings.NewReader("yooo five2")

	mr := io.MultiReader(a, b)

	msg, err := io.ReadAll(mr)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(msg))
}

func six() {
	r := strings.NewReader("yooo six")

	var b bytes.Buffer

	r.WriteTo(&b)

	fmt.Println(b.String())
}
