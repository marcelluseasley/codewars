package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},1))
}

func LongestConsec(strarr []string, k int) string {
	var finalStringSlice []string
	var largestIndex int
	var largestSum int

	// edge cases
	if len(strarr) == 0 || k > len(strarr) || k <= 0 {
		return ""
	}

	for i :=0; i <= len(strarr)-k; i++ {
		sum := 0

		for j := i; j < i+ k; j++ {
			sum += len(strarr[j])
		}
		if sum > largestSum {
			largestSum = sum
			largestIndex = i
		}
	}
	ctr := 0
	for ctr < k {
		finalStringSlice = append(finalStringSlice, strarr[largestIndex+ctr])
		ctr++
	}



	finalString := strings.Join(finalStringSlice,"")

	return finalString

}