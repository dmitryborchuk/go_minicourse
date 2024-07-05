package main

import (
	"fmt"
	"strings"
)

func isVowel(ch int32) bool {
	vowels := "aeiouyAEIOUY"
	return strings.Contains(vowels, (string)(ch))
}

func main() {
	fmt.Println(isVowel('I'))
}
