package main

import "fmt" 

func main(){
	var conferenceName = "Go Conference"
	const confereceTickets = 50
	var remainingTickets uint = 50

	//formated string with place holders ( variable reference)
	fmt.Printf("Welcome to our %v booking applications \n", conferenceName)
	fmt.Println("We have total of", confereceTickets, "tickets and",remainingTickets, "are still available")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets \n", userName,userTickets)
}
