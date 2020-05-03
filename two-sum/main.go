package main

import (
	"fmt"
)

func main() {

	fmt.Println(TwoSum([]int{1234, 5678, 9012}, 14690))
}

func TwoSum(numbers []int, target int) [2]int {
	comps := make(map[int]int)
	results := [2]int{}
	for i, num := range numbers {
		if _, ok := comps[num]; !ok {
			comps[target-num] = i
		} else {
			results[0] = comps[num]
			results[1] = i //append(results, target-num, num)
		}
	}
	return results
}
