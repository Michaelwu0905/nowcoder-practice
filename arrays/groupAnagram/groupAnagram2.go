package main

import (
	"sort"
)

// 辅助函数： 接受一个字符串，返回排序后的字符串
func sortString(w string) string {
	// 1. 将字符串转换为字节切片
	s := []byte(w)

	// 2. 排序
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	// 3. 返回排序好的字符串
	return string(s)
}

// 主函数
func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	for _, str := range strs {
		// 获取排序后的字符串作为key
		sortedStr := sortString(str)

		// 归类
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	// 收集结果
	res := make([][]string, 0, len(mp))
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
