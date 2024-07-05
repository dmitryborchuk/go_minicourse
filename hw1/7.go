package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			fmt.Printf("%d ", i)
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	fmt.Println()
}
