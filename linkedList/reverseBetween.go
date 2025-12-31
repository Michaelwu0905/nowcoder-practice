package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 牛客网，链表内指定区间反转

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	
	// 标准的虚拟头节点起手式
	dummy:=&ListNode{Val:-1}
	dummy.Next=head

	// 找到守卫节点guard，即最后一个不被反转的节点
	pre:=dummy
	for i:=1;i<=m-1;i++{
		pre=pre.Next
	}

	// 定位cur(翻转部分的第一个节点)
	cur:=pre.Next

	// 开始头插法
	for i:=1;i<=n-m;i++{
		temp:=cur.Next		// 先保存cur后面的节点
		cur.Next=temp.Next	// cur越过temp，连接后面
		temp.Next=pre.Next	// temp挤到pre的身后
		pre.Next=temp
	}

	return dummy.Next

}
