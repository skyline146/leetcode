package main

import (
	"fmt"
)

func generateParentheses(n int) []string {
	res := make([]string, 0)

	var makeParentheses func(s string, o, c int)
	makeParentheses = func(s string, o, c int) {
		if o == n && c == n {
			res = append(res, s)
			return
		}

		if o < n {
			makeParentheses(s+"(", o+1, c)
		}

		if c < o {
			makeParentheses(s+")", o, c+1)
		}
	}

	makeParentheses("", 0, 0)
	return res
}

func main() {
	fmt.Println(generateParentheses(4))
}
