package main

import (
	"fmt"
	"slices"
)

func twoSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l]+nums[r] < target {
			l++
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			res = append(res, []int{nums[l], nums[r]})

			sv, sr := nums[l], nums[r]
			for l < r && sv == nums[l] {
				l++
			}

			for l < r && sr == nums[r] {
				r--
			}
		}
	}

	return res
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	slices.Sort(nums)

	for i := range len(nums) - 2 {
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
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{-100000, -99999, 199999}))
	fmt.Println(threeSum([]int{0, 1, 1}))
}
