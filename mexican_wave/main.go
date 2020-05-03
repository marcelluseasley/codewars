package main

import(
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%#v\n",wave(" x yz"))
}

func wave(words string) []string {
	
	result := []string{}
	runes := []rune(words)

	for i, r := range runes {
		if unicode.IsSpace(r) {
			continue
		}

		if unicode.IsLower(r) {			 
			runes[i] = unicode.ToUpper(r)
			result = append(result, string(runes) )
			runes[i] = unicode.ToLower(r)
		}

		

		
	}
	
	return result
}