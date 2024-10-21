package main

import (
	"fmt"
	//"strings"
	"strconv"
	"booking-app/helper"
	"time"
)

var theatreName string = "Go Theatre"
var remainingSeats int = 100 
var bookings = make([]map[string]string,0)
var countId int = 0

// type userData struct{
// 	firstName string
// 	choice bool
// }

// var user = userData{
// 	firstName:firstName
// 	choice:true
// }

// fmt.Println(user.choice)


func main(){

	
	const numberOfSeats uint = 100 //not negative
	var choice string
	var ch string

	greetUsers(theatreName,numberOfSeats,remainingSeats)

	//var bookings=[100]string{"n1","n2","n3"}
	//var bookings [100]string


		getUserData()	

	

	fmt.Println("\nBook Ticket(s)? - [y/n]")
	fmt.Scan(&choice)

	for choice=="y"{
		getUserData()

		fmt.Println("\nBook Ticket? - [y/n]")
		fmt.Scan(&choice)
	}

	fmt.Println("\nWant to Cancel Any Booking??  [y/n]")
	fmt.Scan(&ch)

	if ch=="y"{
		fmt.Println("Enter Id")
		var id int
		fmt.Scan(&id)
		cancelBooking(id)
	}

	printData()

	

}


func greetUsers(theatreName string, numberOfSeats uint, remainingSeats int){
	fmt.Println("\nWelcome to", theatreName, "booking app")
	fmt.Println("We have",numberOfSeats,"seats and remaining seats are",remainingSeats)
	fmt.Println("Book your ticket now\n")
}

func getUserData(){

	var firstName, lastName, email string
	var totalBookingseats int

	fmt.Println("Enter first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter email")
	fmt.Scan(&email)
	fmt.Println("Enter total number of tickets that you want to book ")
	fmt.Scan(&totalBookingseats)

	res := helper.ValidateData(firstName,lastName,email,totalBookingseats,&remainingSeats)

	if res==true{
		proceedForBooking(totalBookingseats,&remainingSeats,email)//pass by reference

		var userData=make(map[string]string)
		userData["Id"]=strconv.Itoa(generateId())
		userData["firstName"]=firstName
		userData["lastName"]=lastName
		userData["email"]=email
		userData["numberOfSeats"]=strconv.Itoa(totalBookingseats)

		bookings=append(bookings,userData)
	}
}


func proceedForBooking(totalBookingseats int,remainingSeats *int,email string){

	fmt.Printf("Thank you for booking %d seats.\n",totalBookingseats)
	fmt.Println("\nProcessing.......\n")
	
	time.Sleep(time.Second*3)

	fmt.Println("Ticket/s are booked.")

	*remainingSeats=*remainingSeats-totalBookingseats

	fmt.Println("After booking remaining seats are: ",*remainingSeats)
	
}

func generateId()int{
	countId+=1
	return countId
}

func cancelBooking(id int) {
	fmt.Println("\nCancelling....\n")
	time.Sleep(time.Second*3)
	fmt.Println("Tickets are Cancelled Successfully!!!\n")
	fmt.Printf("Remaining Tickets are %d",remainingSeats)
	for i, booking := range bookings {
		// Check if the booking ID matches the given ID
		if booking["Id"] == strconv.Itoa(id) {
			// Remove the booking by slicing out the map
			bookings = append(bookings[:i], bookings[i+1:]...)
			//fmt.Println("\nBooking Canceled Successfully")
			return // Exit after deleting
		}
	}
	fmt.Println("Booking ID not found.")
}

func printData(){
	fmt.Println(bookings)
}

