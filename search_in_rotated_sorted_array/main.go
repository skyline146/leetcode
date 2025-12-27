package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2
		x := nums[mid]
		if x == target {
			return mid
		}

		if nums[l] <= x {
			if target < x && nums[l] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > x && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums := []int{2, 4, 5, 6, -1, 0, 1}
	// nums := []int{7, 8, 0, 1, 2, 4, 5, 6}
	fmt.Println(search(nums, 1))
}
