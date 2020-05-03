package main

import (
	"fmt"
)

func main() {
	 
	fmt.Println(FindOutlier([]int{-3, 2, 6, 8, -10, }))
}

func FindOutlier(integers []int) int {
	var evens []int
	var odds []int

	for _, N := range integers {
		
		if N % 2 == 0 {
			evens = append(evens, N)
		} else {
			odds = append(odds, N)
		}
		if len(evens) >= 1 && len(odds) >= 1 {
			if len(evens) > len(odds) {
				return odds[0]
			} else if len(evens) < len(odds) {
				return evens[0]
			}
			
			
		}
	}
	fmt.Println(evens)
	fmt.Println(odds)
	return -1

}