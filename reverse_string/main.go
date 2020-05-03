/*

Given a string N, reverse the string and return it

N -> "orchestration"  ReverseString(N) -> "noitartsehcro"


*/
package main


import (
	"fmt"
	
)
func main() {

	N := "orchestration"
	fmt.Println(ReverseString(N))
}

func ReverseString(N string) string {

	runes := []rune(N)
	
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}