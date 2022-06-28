package main

import (
	"errors"
	"fmt"
	"io"
)

type myType []byte

func (m myType) Read(p []byte) (n int, err error) {
	for range p {
		n++
	}

	return n, io.EOF
}

func (m myType) Write(p []byte) (n int, err error) {
	for range p {
		n++
	}
	if n != len(string(p)) {
		err = errors.New("Ошибка считывания")
	}
	return n, err
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func main() {
	m := myType{}

	io.WriteString(m, "Привет")
	fmt.Println(m)

	myStr, _ := io.ReadAll(m)

	fmt.Println(string(myStr))
}
