package main

import "fmt"

type Address struct {
	Name string

	city string

	Pincode int
}

func main() {

	var a Address

	fmt.Println(a)

	a1 := Address{"Sharan", "tirunelveli", 627007}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "vengy", city: "chennai",

		Pincode: 6000027}

	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "madurai"}

	fmt.Println("Address3: ", a3)
}
