package main

import (
	"fmt"
	"slices"
)

func twoSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	m := make(map[int][]int)

	for i, v := range nums {
		if a, ok := m[target-v]; ok {
			for _, j := range a {
				minv, maxv := min(nums[i], nums[j]), max(nums[i], nums[j])

				res = append(res, []int{minv, maxv})
			}
		}

		m[v] = append(m[v], i)
	}

	return res
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		v := nums[i]

		if i != 0 && v == nums[i-1] {
			continue
		}

		for _, k := range twoSum(nums[i+1:], -v) {
			res = append(res, append(k, v))
		}
	}

	return res
}

func main() {
	// fmt.Println(twoSum([]int{0, 1, 2, -1, 1, -4}, 1))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{-100000, -99999, 199999}))
	// fmt.Println(threeSum([]int{0, 0, 0}))
}
