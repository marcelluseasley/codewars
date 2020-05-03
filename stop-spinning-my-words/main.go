package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SpinWords("This is another test"))
}

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	var strSlice []string

	for _, word := range words {
		if len(word) < 5 {
			strSlice = append(strSlice, word)
			continue
		}
		strSlice = append(strSlice, flip(word))
	}
	return strings.Join(strSlice, " ")
}

func flip(str string) string {
	word := strings.Trim(str," ")
	runes := []rune(word)

	for i, j := 0, len(runes)-1 ; i< len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}