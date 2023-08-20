package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const confereceTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//formated string with place holders ( variable reference)
	fmt.Printf("Welcome to our %v booking applications \n", conferenceName)
	fmt.Println("We have total of", confereceTickets, "tickets and", remainingTickets, "are still available")

	var bookings []string

	// ask user for their name

	//fmt.Println(remainingTickets)
	//fmt.Println(&remainingTickets)

	//ask for name:
	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name :")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email :")
	fmt.Scan(&email)

	fmt.Println("Enter no.of tickets :")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v ", remainingTickets, conferenceName)
}
