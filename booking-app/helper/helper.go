package helper

import (
	"strings"
	"fmt"
)

func ValidateData(firstName string,lastName string,email string,totalBookingseats int,remainingSeats *int) bool{

	isValidName := len(firstName)>=2 && len(lastName)>=2
	isValidEmail := strings.Contains(email,"@")
	isValidNumber := totalBookingseats <= 0 || totalBookingseats > *remainingSeats

	if !isValidEmail{
		fmt.Println("Invalid Email Format\n")
	}else if !isValidName{
		fmt.Println("Invalid Name Format\n")
	}else if isValidNumber{
		fmt.Printf("The Seats that you want to book must be positive and less than %d seats",remainingSeats)
	}else{
		return true
	}

	return false
}


//go run main.go helper.go  or   go run .