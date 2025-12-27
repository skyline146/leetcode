package main

import (
	"fmt"
	"slices"
)

func sortString(word string) string {
	runes := []rune(word)
	slices.Sort(runes)
	return string(runes)
}

func removeAnagrams(words []string) []string {
	res := []string{words[0]}
	prevSorted := sortString(res[0])

	for _, word := range words[1:] {
		sortedWord := sortString(word)
		if prevSorted == sortedWord {
			continue
		}

		prevSorted = sortedWord
		res = append(res, word)
	}

	return res
}

func main() {
	fmt.Println(removeAnagrams([]string{"truqjvrb"}))
}
