package main

import "fmt"
func main() {
	fmt.Println("THE ARRAYS")
	var myarr[3]string
	myarr[0] = "DATA 1"
	myarr[1] = "DATA 2"
	myarr[2] = "DATA 3"
	fmt.Println("the elements are",myarr)
	fmt.Println("the first element",myarr[0])
	fmt.Println("the second element",myarr[1])
	fmt.Println("the third element",myarr[2])
}