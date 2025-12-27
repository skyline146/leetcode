package main

import (
	"fmt"
	"strings"
)

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	pr := []rune(p)
	sr := []rune(s)
	i := 0
	starIdx := -1
	matchIdx := 0

	for j := 0; j < len(sr); {
		if i < len(pr) && (sr[j] == pr[i] || pr[i] == '?') {
			i++
			j++

			continue
		}

		if i < len(pr) && pr[i] == '*' {
			starIdx = i
			matchIdx = j
			i++

			continue
		}

		if starIdx == -1 {
			return false
		} else {
			i = starIdx + 1
			matchIdx++
			j = matchIdx
		}
	}

	return len(pr[i:]) == strings.Count(string(pr[i:]), "*")
}

func main() {
	fmt.Println(isMatch("abcxabczzzde", "*xabc???de"))
}

// slow
func isMatchRec(s string, p string) bool {
	if s == "" {
		if p == "" || len(p) == strings.Count(p, "*") {
			return true
		}

		return false
	}

	if p == "" {
		return s == ""
	}

	if s[0] == p[0] || p[0] == '?' {
		return isMatchRec(s[1:], p[1:])
	}

	if p[0] == '*' {
		return isMatchRec(s, p[1:]) || isMatchRec(s[1:], p)
	}

	return false
}
