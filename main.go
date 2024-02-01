package main

import (
	"fmt"
	"strconv"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]map[string]string , 0)

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)
			firstName:= getFirstNames()

			fmt.Printf("The first names %v\n", firstName)
			if remainingTickets == 0 {
				break
			}
		} else {
			 if !isValidName {
				fmt.Println("Check your first or last name its too short")
			 } 

			 if !isValidEmail {
				fmt.Println("Enter a valid email")
			 }

			 if !isValidTicketNumber {
				fmt.Println("Enter a valid ticket number")
			 }

			 continue
		}
	}
	

}


func greetUsers() {
	fmt.Printf("Welcome to our %v Booking Application\n You can Book Your Conference Tickets from Here.\nWe Have Total of %v Tickets and still % v Tickets are available.\n", conferenceName, conferenceTickets, remainingTickets)
}

func getUserInputs() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter your ticket amount:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames:= []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])

	}

	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber


}


func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets
	
	var user = make(map[string]string)
	user["firstName"] = firstName
	user["lastName"] = lastName
	user["email"] = email
	user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)


	bookings = append(bookings, user)


	fmt.Printf("Thank you! %v %v for booking %v tickets. You'll receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v ticket(s) still available\n", remainingTickets)
	fmt.Printf("These are all our bookings: %v\n", bookings)

}