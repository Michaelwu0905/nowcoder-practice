package main

// 力扣：最长不重复字串

func lengthOfLongestSubstring(s string) int{
	// 边界条件，s为空则返回
	if len(s)==0{
		return 0
	}

	// 准备变量
	var freq [128]int
	left:=0
	maxLen:=0

	for right:=0;right<len(s);right++{
		char:=s[right]
		freq[char]++

		// 若s[right]的数量大于1，说明就是有重复的子串
		// left往右缩，知道字符不重复
		for freq[char]>1{
			freq[s[left]]--
			left++
		}

		// 此时窗口合法，计算当前长度然后更新maxlen
		// 长度为right-left+1
		currentLen:=right-left+1
		if currentLen>maxLen{
			maxLen=currentLen
		}
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int{
	if len(s)==0{
		return 0
	}

	// 1. 显示地转换成byte切片
	// byte 即uint8，范围0-255，需要256大小的数组
	sBytes:=[]byte(s)
	var freq [256]int
	left:=0
	maxLen:=0

	// 2. 遍历byte切片
	for right,b := range sBytes{
		freq[b]++

		// 窗口收缩
		for freq[b]>1{
			// 移除左边界的byte
			leftByte:=sBytes[left]
			freq[leftByte]--
			left++
		}

		currentLen:=right-left+1
		if currentLen>maxLen{
			maxLen=currentLen
		}
	}
	return maxLen
}

func lengthOfLongestSubstring3(s string) int{
	if len(s)==0{		// 采用map，最直观的方法
		return 0
	}

	// 显式定义，key是byte，value是出现次数
	freq:=make(map[byte]int)

	left:=0
	maxLen:=0

	for right:=0;right<len(s);right++{
		char:=s[right]
		freq[char]++

		for freq[char]>1{
			leftByte:=s[left]
			freq[leftByte]--
			left++
		}

		currentLen:=right-left+1
		if currentLen>maxLen{
			maxLen=currentLen
		}
	}
	return maxLen
}