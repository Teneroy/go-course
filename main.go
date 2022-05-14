package main

import (
	"fmt"
	"go-course/helper"
	"strconv"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) //Slice

// struct does not support inheritance
//type UserData struct {
//	firstName string
//	lastName string
//	email string
//	numberOfTickets uint
//}

func main() {
	fmt.Printf("conferenceTickets is of type %T, remainingTickets is of type %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		userName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidTicketNumber && isValidEmail {

			bookTicket(userTickets, userName, lastName, email)
			fmt.Printf("First names of our bookings are: %v\n", getFirstNames())

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

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
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

func bookTicket(userTickets uint, userName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = userName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v booked %v tickets. You will receive a confirmation on %v\n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
