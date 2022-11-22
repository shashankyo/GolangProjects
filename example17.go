package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 67, 8, 9}

	b := a[:]
	c := a[3:]
	d := a[3:6]
	e := a[:6]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
