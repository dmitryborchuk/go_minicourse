package main

import "fmt"

func reverse(s string) string {
	str := []rune(s)
	i := 0
	j := len(str) - 1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return (string)(str)
}

func main() {
	fmt.Println(reverse("accaabacda"))
}
