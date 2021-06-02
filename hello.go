package main

import "fmt"
type author struct{
	name  string
	age int
	code int
	salary int
}
func (a author) show() {
	fmt.Println("name:", a.name)
	fmt.Println("age:", a.age)
	fmt.Println("code is :", a.code)
	fmt.Println("salary is :", a.salary)
}
func main() {
	res := author{
	name : "sharan",
	age : "21",
	code : "TR/21/2021",
	salary : "10000," 
}
res.show()
}