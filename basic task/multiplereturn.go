package main

import "fmt"

func vals() (int, int) {
    return 350,700
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
}