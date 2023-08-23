package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			bookTicket(userTickets, bookings, firstName, lastName, email, conferenceName)

			// function print first names
			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings : %v \n ", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end the program
				fmt.Println("Our conference is booked out. come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email you entered is not valid")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			//fmt.Printf("We only have %v tickets remaining , so you can't book %v tickets\n", remainingTickets, userTickets)
			fmt.Println("Your input data is invalid, try again")
			continue
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking applications \n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name :")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last name :")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email :")
	fmt.Scanln(&email)

	fmt.Println("Enter no.of tickets :")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - uint(userTickets)

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v ", remainingTickets, conferenceName)
}
