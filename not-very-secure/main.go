package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(alphanumeric("hello world_"))
}

func alphanumeric(str string) bool {
	if len(str) < 1 {
		return false
	}
	matches, _ := regexp.Match(`^[a-zA-Z0-9]+$`,[]byte(str))
	return matches
}