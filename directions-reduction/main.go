package main

import "fmt"

func main() {
	fmt.Println(DirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}))
}

func DirReduc(arr []string) []string {
	dirMap := map[string]string{"NORTH":"SOUTH", "SOUTH":"NORTH", "EAST":"WEST","WEST":"EAST"}

	var directions Stack
	for _, a := range arr {
		if !directions.IsEmpty() && dirMap[a] == directions[len(directions)-1]{
			directions.Pop()
		} else {
			directions.Push(a)
		}

	}

	return directions
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop()  {
	if s.IsEmpty() {
		return
	} else {
		index := len(*s) - 1
		*s = (*s)[:index]
	}
}