package main

import "fmt"

func main() {

	var year int = 2000
	if year%4 == 0 && year%100 == 0 && year%400 == 0 {
		fmt.Println(year, "is Leap Year")
	} else {
		fmt.Println(year, "is Not leap year")
	}
}
