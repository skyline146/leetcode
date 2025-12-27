package main

import "fmt"

// my first solution
func lengthOfLongestSubstring2(s string) int {
	m := make(map[byte]int)

	max, k := 0, 0
	for i := range len(s) {
		c := s[i]
		x, ok := m[c]
		if ok {
			if k > max {
				max = k
			}

			between := i - x

			if between == 1 {
				clear(m)

				k = 1
				m[c] = i

				continue
			}

			if k > between-1 {
				k = between - 1
			}
		}

		k++
		m[c] = i
	}

	if k > max {
		return k
	}

	return max
}

// optimized
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)

	res, left := 0, -1
	for i := range len(s) {
		c := s[i]
		if x, ok := m[c]; ok && x > left {
			left = x
		}

		res = max(i-left, res)

		m[c] = i
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("ckilbkd"))  // 5
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // 3
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring("cdd"))      // 2
	fmt.Println(lengthOfLongestSubstring("abba"))     // 2
	fmt.Println(lengthOfLongestSubstring("wobgrovw")) // 6
}
