package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	tmp, revX := x, 0

	for tmp != 0 {
		revX, tmp = revX*10+tmp%10, tmp/10
	}

	return x == revX
}

func main() {
	fmt.Println(isPalindrome(1221))
}
