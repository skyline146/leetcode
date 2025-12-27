package main

import (
	"fmt"
)

func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}

	prev := grayCode(n - 1)
	curr := make([]int, 0, len(prev))

	for i := len(prev) - 1; i >= 0; i-- {
		curr = append(curr, prev[i]|1<<(n-1))
	}

	return append(prev, curr...)
}

func main() {
	fmt.Println(grayCode(4))
}
