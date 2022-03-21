package utilities

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketRequest := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketRequest
}
