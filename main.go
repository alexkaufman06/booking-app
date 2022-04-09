package main

import (
	"booking-app/utilities"
	"fmt"
	"sync"
	"time"
)

// https://www.youtube.com/watch?v=yyUHQIec83I 3:15:20

// package level variables
const venueTickets = 50
const venue = "Mississippi Studios"

var bookings = make([]UserData, 0)
var remainingTickets uint = 50 // unsigned integers can't be negative

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	userFirstName, userLastName, userEmail, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketRequest := utilities.ValidateUserInput(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketRequest {
		bookTicket(userTickets, userFirstName, userLastName, userEmail)
		wg.Add(1)
		go sendTicket(userTickets, userFirstName, userLastName, userEmail)
		printFirstNames()

		if remainingTickets == 0 {
			fmt.Printf("Sorry, this show is sold out. Come back next time.")
		}
	} else {
		fmt.Println("Your input is invalid, please try again:")
		if !isValidName {
			fmt.Println("  * First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("  * Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketRequest {
			fmt.Println("  * The number of tickets you requested is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v concert ticket booking application!\n", venue)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", venueTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	fmt.Println("What is your first name?")
	fmt.Scan(&userFirstName)

	fmt.Println("What is your last name?")
	fmt.Scan(&userLastName)

	fmt.Println("What is your email?")
	fmt.Scan(&userEmail)

	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, venue)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#########################")
	wg.Done()
}
