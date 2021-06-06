package main
import (
    "fmt"
)
func main() {

    days := 2
    fmt.Print("Write ", days, " as ")
    switch days {
    case 1:
        fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	}
}
