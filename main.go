package main

import "fmt"

func main() {
	fmt.Println(addArray([]int{1, 2, 3}, []int{4, 5, 6}))
}

func addArray(left, right []int) []int {
	return append(left, right...)
}
