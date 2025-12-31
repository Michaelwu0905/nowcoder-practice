package main
import "fmt"
func reverseString(s []byte){

	if len(s)==0{
		return
	}

	left:=0
	right:=len(s)-1

	for left<right{
		s[left],s[right]=s[right],s[left]

		left++
		right--
	}

	// 原地修改，不用返回值
}