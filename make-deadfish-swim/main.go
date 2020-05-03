package main

import (
	"fmt"
)

func main() {
	fmt.Println(Parse("isoisoiso"))

}

func Parse(data string) []int {
	s := []int{}
	runes := []rune(data)

	num := 0

	for _, r := range runes {
		switch r {
		case 'i':
			num++
		case 'd':
			num--
		case 's':
			num *= num
		case 'o':
			s = append(s, num)
		}
	}
	return s
}
