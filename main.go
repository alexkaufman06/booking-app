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

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
