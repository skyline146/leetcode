package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	revX, tmp := int64(0), int64(x)
	for tmp != 0 {
		revX, tmp = revX*10+tmp%10, tmp/10

		if revX < math.MinInt32 || revX > math.MaxInt32 {
			return 0
		}
	}

	return int(revX)
}

func main() {
	fmt.Println(reverse(-123))
}
