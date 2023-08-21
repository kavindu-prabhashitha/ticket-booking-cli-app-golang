package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//formated string with place holders ( variable reference)
	fmt.Printf("Welcome to our %v booking applications \n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

	var bookings []string

	for {
		fmt.Println("Enter your first name :")
		fmt.Scanln(&firstName)

		fmt.Println("Enter your last name :")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your email :")
		fmt.Scanln(&email)

		fmt.Println("Enter no.of tickets :")
		fmt.Scanln(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v ", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings : %v \n ", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end the program
				fmt.Println("Our conference is booked out. come back next year")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining , so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

	}

	// ask user for their name

	//fmt.Println(remainingTickets)
	//fmt.Println(&remainingTickets)

	//ask for name:

}
