package main

import "fmt"

func main() {
	var students [3]string
	fmt.Println("students:%v\n", students)

	students[0] = "lisa"
	fmt.Println("students:%v\n", students)

}
