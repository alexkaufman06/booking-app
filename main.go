package main

import (
	"fmt"
	"strings"
)

// https://www.youtube.com/watch?v=yyUHQIec83I 1:54:13

func main() {
	const venueTickets = 50
	venue := "Mississippi Studios"
	bookings := []string{}
	var remainingTickets uint = 50 // unsigned integers can't be negative

	fmt.Printf("Welcome to our %v concert ticket booking application!\n", venue)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", venueTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
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

		isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketRequest := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketRequest {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userFirstName+" "+userLastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userFirstName, userLastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, venue)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Sorry, this show is sold out. Come back next time.")
				break
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
	}
}
