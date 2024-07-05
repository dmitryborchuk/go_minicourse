package main

import "fmt"

func getArraysSum(arr []int) int {
	ans := 0
	for _, val := range arr {
		ans += val
	}
	return ans
}

func main() {
	test := []int{1, 5, 9, -1, 4}
	fmt.Println(getArraysSum(test))
}
