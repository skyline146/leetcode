package main

import (
	"fmt"
)

// func twoSum(nums []int, target int) []int {
// 	m := make(map[int]int)
// 	for i, num := range nums {
// 		if j, ok := m[target-num]; ok {
// 			return []int{i, j}
// 		}
// 		m[num] = i
// 	}

// 	return []int{}
// }

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	isInit := false

	valToIdx := make(map[int][]int)
	for i, v := range nums {
		if len(valToIdx[v]) < 2 {
			valToIdx[v] = append(valToIdx[v], i)
		}
	}

	for key, idx1Slice := range valToIdx {
		if idx2Slice, ok := valToIdx[target-key]; ok {
			idx1, idx2 := idx1Slice[0], idx2Slice[0]

			if idx1 == idx2 {
				if len(idx2Slice) == 1 {
					continue
				}

				idx2 = idx2Slice[1]
			}

			if idx1+idx2 < result[0]+result[1] || !isInit {
				result[0], result[1] = idx1, idx2
				isInit = true
			}
		}
	}

	return result
}

func main() {
	// fmt.Println(twoSum([]int{-2, 8, -10, 1, -7, 11, 15, -17}, -9))
	fmt.Println(twoSum([]int{1, 1, 1, 1}, 2))
}
