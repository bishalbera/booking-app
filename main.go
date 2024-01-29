package main
import "fmt"


func main() {

	conferenceName:= "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 30
	fmt.Printf("Welcome to our %v Booking Application\n", conferenceName)
	fmt.Println("You can Book Your Conference Tickets from Here.")
	fmt.Printf("We Have Total of %v Tickets and still % v Tickets are available.\n", conferenceTickets, remainingTickets )

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter your ticket amount")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you! %v %v for booking %v tickets. You'll receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v ticket(s) still available\n", remainingTickets)

}