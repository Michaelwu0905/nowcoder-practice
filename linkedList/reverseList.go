package main

// 牛客：翻转链表

type ListNode struct{
	Val int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode{
	
	var pre *ListNode=nil
	cur:=head

	for cur!=nil{
		temp:=cur.Next
		cur.Next=pre
		pre=cur
		cur=temp
	}
	return pre
}