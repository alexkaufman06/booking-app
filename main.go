package main

import "fmt"

// https://www.youtube.com/watch?v=yyUHQIec83I 45:30

func main() {
	const venueTickets = 50
	venue := "Mississippi Studios"
	var remainingTickets uint = 50 // unsigned integers can't be negative

	fmt.Printf("venueTickets is %T, remainingTickets is %T, venue is %T\n", venueTickets, remainingTickets, venue)

	fmt.Printf("Welcome to our %v concert ticket booking application!\n", venue)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", venueTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets int

	fmt.Println("What is your first name?")
	fmt.Scan(&userFirstName)

	fmt.Println("What is your last name?")
	fmt.Scan(&userLastName)

	fmt.Println("What is your email?")
	fmt.Scan(&userEmail)

	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
}
