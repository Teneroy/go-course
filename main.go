package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is of type %T, remainingTickets is of type %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string //Slice

	for remainingTickets > 0 && len(bookings) < 50 {
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

		var isValidName = len(userName) >= 2 && len(lastName) >= 2
		var isValidEmail = strings.Contains(email, "@")
		var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidTicketNumber && isValidEmail {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName+" "+lastName)

			fmt.Printf("Thank you %v %v booked %v tickets. You will receive a confirmation on %v\n", userName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First names of our bookings are: %v\n", firstNames)

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
