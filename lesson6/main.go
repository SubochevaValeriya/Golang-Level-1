package main

import (
	"fmt"
	"io"
)

type myType []byte

func (m *myType) Read(p []byte) (n int, err error) {

	for i, value := range *m {
		p[i] = value
	}

	n = len(*m)
	return n, io.EOF
}

func (m *myType) Write(p []byte) (n int, err error) {
	*m = p

	n = len(*m)

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

	io.WriteString(&m, "Привет")

	fmt.Println(string(m))

	myStr, _ := io.ReadAll(&m)

	fmt.Println(string(myStr))
}
