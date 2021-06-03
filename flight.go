package main

import "fmt"

func main() {
	fmt.Println("flight ticket booking ")
	fmt.Println("enter your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("hi", name)
	fmt.Println("ENTER MAIL ID")
	var mailid string
	fmt.Scanln(&mailid)
	fmt.Println("from")
	var from string
	fmt.Scanln(&from)
	fmt.Println("to")
	var to string
	fmt.Scanln(&to)
	fmt.Println("PLEASE CHOOSE YOUR TRAVEL DATES WITHIN THIS MONTH")
	var date int
	fmt.Scanln(&date)
	if date == 0 {
		fmt.Println("sorry it is not available")
	} else {
		fmt.Println("COOL..! FASTEN YOUR SEATBELTS")
	}
	var flights [3]string
	flights[0] = "SPICE JET"
	flights[1] = "KING FISHER"
	flights[2] = "DECCAN AIR"
	fmt.Println("THE AVAILABLE FLIGHTS ON YOUR DATE ARE", flights)
	fmt.Println("CHOOSE YOUR FLIGHT")
	fmt.Println("choose from 1 2 3")
	var choose int
	fmt.Scanln(&choose)
	if choose == 1 {
		fmt.Println("WELCOME TO SPICE JET")
	}

	if choose == 2 {
		fmt.Println("WELCOME TO KING FISHER ")
	}
	if choose == 3 {
		fmt.Println("WELCOME TO DECCAN AIR")
	}
	fmt.Println("PLEASE MENTION NUMBER OF SEATS")
	var seat int
	fmt.Scanln(&seat)
	fmt.Println("PLEASE WAIT WE ARE PROCESSING YOUR REQUEST")
	if seat >= 4 {
		fmt.Println("SORRY ONLY 3 TICKETS CAN BE PROCESSED AT A TIME")
	} else {
		fmt.Println("REDIRECTED TO PAYMENT PAGE")
	}
	fmt.Println("ONE TICKET PRICE IS 4000 FOR ALL.")
	fmt.Println("CONFIRM NUMBER OF SEATS")
	var tickets int
	fmt.Scanln(&tickets)
	if tickets == 1 {
		fmt.Println("pay rs 4000")
	}
	if tickets == 2 {
		fmt.Println("pay rs 8000")
	}
	if tickets == 3 {
		fmt.Println("pay rs 12000")
	}
	fmt.Println(" PLEASE CHOOSE YOUR PAYMENT METHOD")
	var payment [3]string
	payment[0] = "DEBIT CARD"
	payment[1] = "CREDIT CARD"
	payment[2] = "UPI ID"
	fmt.Println(payment)
	var number int
	fmt.Scanln(&number)
	fmt.Println("PLEASE WAIT UNTIL WE PROCESS YOUR REQUEST")
	if number == 1 {
		fmt.Println("DEBIT CARD ACCEPTED")
		fmt.Println(" ENTER CARD NO ")
		var cardno int
		fmt.Scanln(&cardno)
		fmt.Println("ENTER EXPIRY DATE")
		var expdate int
		fmt.Scanln(&expdate)
		fmt.Println("ENTER CVV")
		var cvv int
		fmt.Scanln(&cvv)
		fmt.Println("PLEASE WAIT")

	}
	if number == 2 {
		fmt.Println("CREDIT CARD ACCEPTED")
		var cardnu int
		fmt.Scanln(&cardnu)
		fmt.Println("ENTER EXPIRY DATE")
		var expdatu int
		fmt.Scanln(&expdatu)
		fmt.Println("PLEASE WAIT")

	}
	if number == 3 {
		fmt.Println("UPI ACCEPTED")
		fmt.Println("ENTER UPI ID")
		var upi string
		fmt.Scanln(&upi)
		fmt.Println("PLEASE WAIT")
	}
	fmt.Println("BOOKING SUCCESSFULL")
	fmt.Println("DETAILS WILL BE SEND TO YOUR REGISTERD MAIL ID")
	fmt.Println("THANK YOU..!.. HAVE A NICE JOURNEY")

}
