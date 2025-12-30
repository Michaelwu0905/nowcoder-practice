package main

import "fmt"

// 力扣：字母异位词分组
// 利用go能直接比较数组的特性，直接拿数组当作键值

func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, char := range str {
			// 统计每个字符出现的频次
			count[char-'a']++
		}
		// 将统计数组记为key，原始字符串加入value
		mp[count] = append(mp[count], str)
	}

	res := make([][]string, 0, len(mp))
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func main() {
	strs := []string{"abc", "cba", "kka", "akk", "fcku"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
