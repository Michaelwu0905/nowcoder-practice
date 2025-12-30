package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := m[complement]; found {
			return []int{index, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
