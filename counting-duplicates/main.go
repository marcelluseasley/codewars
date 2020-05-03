package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.ToLower("aabBcde")
	fmt.Println(duplicate_count(s))
}

package kata

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) int {
	//your code goes here

	s1 = strings.ToLower(s1)

	count := 0

	runeMap := make(map[rune]int)
	runes := []rune(s1)

	for _, r := range runes {
		runeMap[r]++
	}

	for _, v := range runeMap {
		if v >= 2 {
			count++
		}
	}
	fmt.Printf("%s = %d\n", s1, count)
	return count 
}
