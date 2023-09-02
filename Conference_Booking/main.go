package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	var tickets = 100
	var firstName string
	var lastName string
	var emailAddress string
	var totalTickets uint
	var remainingTickets uint = 100

	fmt.Printf("Welcome to %v booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaialable\n", tickets, tickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("Enter your First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email Address")
	fmt.Scan(&emailAddress)

	fmt.Println("Enter the tickets you want to Purchase")
	fmt.Scan(&totalTickets)

	remainingTickets = remainingTickets - totalTickets
	fmt.Printf("Thank you %v%v for purchasing %v tickets, You will receive a confirmation message to the registered %v address\n", firstName, lastName, totalTickets, emailAddress)

	fmt.Println(remainingTickets, "Tickets Are Remaining")

}
