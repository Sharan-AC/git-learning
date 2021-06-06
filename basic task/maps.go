package main
import "fmt"
func main() {
	days := make(map[string]int)
	days["monday"] = 1
	days["tuesday"] = 2
	days["wednesday"] = 3
	days["thursday"] = 4
	days["friday"] = 5
	fmt.Println(days)
}