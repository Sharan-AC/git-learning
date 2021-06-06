package main
import "fmt"

func main(){
	names := [5]string{"hi","hello","welcome","am","sharan"}
	var slice[]string = names[1:4]
	fmt.Println(slice)
}