package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
    for _, words := range words {
        result[words]++
            
    }
    return result
}


func main() {
	wc.Test(WordCount)
}
