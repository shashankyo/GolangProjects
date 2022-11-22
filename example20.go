package main

import "fmt"

// func main() {
// 	a := []int{}
// 	fmt.Println(a)
// 	fmt.Printf("length:%v\n", len(a))
// 	fmt.Println("capacity :%v\n", cap(a))

// 	a = append(a, 1)
// 	fmt.Println(a)
// 	fmt.Printf("length:%v\n", len(a))
// 	fmt.Printf("capacity:%v\n", cap(a))
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	b := a[1:]
// 	fmt.Println(b)
// }

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := append(a[:2], a[3:]...)
	fmt.Println(b)
	fmt.Println(a)
}
