package main 

import "fmt"
func area(length, width int)int{ 

    Areas := length* width 

    return Areas 
} 
 
func main() { 

   fmt.Printf("Area of rectangle is : %d", area(100, 10)) 
}