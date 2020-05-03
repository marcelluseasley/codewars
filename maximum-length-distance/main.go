package main

import (
	"fmt"
	
)

func main(){
	var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	
	fmt.Println(MxDifLg(s1,s2))
}

func MxDifLg(a1 []string, a2 []string) int {
	// your code
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	
	a1Short := len(a1[0])
	a1Long := len(a1[0])

	a2Short := len(a2[0])
	a2Long := len(a2[0])

	for _, a := range a1 {
		if a1Short > len(a) {
			a1Short = len(a)
		}
		if a1Long < len(a) {
			a1Long = len(a)
		}
	}

	for _, a := range a2 {
		if a2Short > len(a) {
			a2Short = len(a)
		}
		if a2Long < len(a) {
			a2Long = len(a)
		}
	}

	max := a1Long - a2Short
	if a2Long - a1Short > max {
		max = a2Long - a1Short
	}
	return max
}