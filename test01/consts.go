package main

import "fmt"

type demo interface {
	SayHi(string)
	Error() error
}

func main() {
	fmt.Println("hello world")
}
