package main

import "fmt"

var a int = 10
var b int = 20

func add() {
	res := (a + b)
	fmt.Println("addition value is", res)
}
func div() {
	res := (a / b)
	fmt.Println("division value is", res)
}
func sub() {
	res := (a - b)
	fmt.Println("substraction value is", res)
}
func mul() {
	res := (a * b)
	fmt.Println("multiplication value is", res)
}
func main() {
	add()
	mul()
	div()
	sub()
}
