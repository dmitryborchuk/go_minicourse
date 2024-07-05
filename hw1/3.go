package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x%2 == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}
