package main

import "fmt"

// leetcode: 两数之和
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		need := target - nums[i]

		// 用go语言的ok模式检查key是否真的存在
		if index, ok := seen[need]; ok {
			return []int{index, i}
		}

		seen[nums[i]] = i
	}

	return nil
}

// leetcode: 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	charPattern := make(map[[26]byte][]string) //

	for _, str := range strs {
		var cnt [26]byte

		for i := 0; i < len(str); i++ {
			cnt[str[i]-'a']++
		}

		charPattern[cnt] = append(charPattern[cnt], str)
	}

	res := make([][]string, 0, len(charPattern))
	for _, v := range charPattern {
		res = append(res, v)
	}

	return res
}

// leetcode: 最长连续序列
func longestConsecutive(nums []int) int {
	
	// 1. 建立哈希集合
	seen:=make(map[int]bool)
	for _,num:=range nums{
		seen[num]=true
	}

	maxLen:=0

	// 遍历集合中的每个数字
	for num:=range seen{
		if !seen[num-1]{
			currentNum:=num
			currentLen:=1

			// 往后数，看num+1，num+2是否存在
			for seen[currentNum+1]{
				currentNum++
				currentLen++
			}

			// 更新全局最大长度
			if currentLen>maxLen{
				maxLen=currentLen
			}
		}
	}
	return maxLen
}

func main() {

	fmt.Println("leetcode: 两数之和 测试样例")
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)

	fmt.Println("leetcode: 字母异位词分组 测试样例")
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res2 := groupAnagrams(strs)
	fmt.Println(res2)

}
