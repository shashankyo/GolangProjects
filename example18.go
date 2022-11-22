package main

import "fmt"

func main() {
	a := make([]int, 8)
	fmt.Println(a)
	fmt.Printf("lemgth:%v\n", cap(a))
	fmt.Printf("capacity :%v\n", cap(a))
}
