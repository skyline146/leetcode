package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1

	k := max(i, j) + 1
	res := make([]byte, k+1) // +1 -> because 1 additional space for 0 index if overflow
	var carry byte

	for k > 0 {
		x := carry

		if i >= 0 {
			x += a[i] & 1
			i--
		}

		if j >= 0 {
			x += b[j] & 1
			j--
		}

		carry = x >> 1
		res[k] = x&1 + '0'
		k--
	}

	if carry == 1 {
		res[0] = '1'
	} else {
		res = res[1:]
	}

	return string(res)
}

func main() {
	fmt.Println(addBinary("1000", "1101"))
}
