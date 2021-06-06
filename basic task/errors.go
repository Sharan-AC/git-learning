package main
import "fmt"
type Myerror struct{}
func (myerr *Myerror) Error()string {
	return "Sorry error..!"
}
func main(){
	myerr := &Myerror{}
	fmt.Println(myerr)
}