package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Println("%v\n", a)
	fmt.Println("%v\n", b)
	fmt.Println("%v\n", c)

}
