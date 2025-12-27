package main

import "fmt"

func getDescentPeriods(prices []int) int64 {
	var count int64 = 1
	prev := prices[0]
	sequenceLen := 0

	for i := 1; i < len(prices); i++ {
		count++

		if prev-prices[i] == 1 {
			count++
			sequenceLen++
		} else {
			sequenceLen = 0
		}

		if sequenceLen > 1 {
			count += int64(sequenceLen) - 1
		}

		prev = prices[i]
	}

	return count
}

func main() {
	fmt.Println(getDescentPeriods([]int{6, 5, 4, 3, 2, 1, 4, 3}))
}
