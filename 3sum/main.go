package main

import "fmt"

func twoSum(nums []int, target int) [][2]int {
	res := make([][2]int, 0)

	m := make(map[int][]int)
	for i, v := range nums {
		if a, ok := m[target-v]; ok {
			for _, j := range a {
				res = append(res, [2]int{j, i})
			}
		}

		m[v] = append(m[v], i)
	}

	return res
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	m := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = make([]int, 0)
		}

		m[nums[i]] = append(m[nums[i]], i)
	}

	return res
}

func main() {
	fmt.Println(twoSum([]int{-1, 0, 1, 2, -1, -4, 1, 1, 0, 1}, 2))
}
