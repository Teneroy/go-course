package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string //Slice

	fmt.Printf("conferenceTickets is of type %T, remainingTickets is of type %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {

		userName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidTicketNumber && isValidEmail {

			bookTicket(remainingTickets, userTickets, bookings, userName, lastName, email, conferenceName)
			fmt.Printf("First names of our bookings are: %v\n", getFirstNames(bookings))

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email has no @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
		}
	}
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func greetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func validateUserInput(userName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(userName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var userName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Print("Enter your first name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return userName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, userName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userName+" "+lastName)

	fmt.Printf("Thank you %v %v booked %v tickets. You will receive a confirmation on %v\n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
