package main

import "fmt"

type Stack[T any] struct { // 泛型栈，可以存放任何类型的数据
	items []T
}

// NewStack 创建一个新栈
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// 核心方法
// 入栈
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop出栈
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T // 声明T类型的0值
		return zero, false
	}

	// 获取栈顶索引
	index := len(s.items) - 1
	// 取出元素
	item := s.items[index]

	// 细节：为防止内存泄漏（T是指针的情况下），最好显式地置空
	var zero T
	s.items[index] = zero

	// 切片缩容(去掉最后一个元素)
	s.items = s.items[:index]

	return item, true
}

// Peek 仅查看栈顶元素但不删除
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	// 直接返回最后一个元素
	return s.items[len(s.items)-1], true
}

// 辅助方法
// 判空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 获取大小
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Print 打印栈道内容
func (s *Stack[T]) Print() {
	fmt.Printf("Current elements in Stack: %v, Size: %d\n", s.items, s.Size())
}

// 主函数测试
func main() {
	fmt.Println("---测试整型栈(Integer Stack)---")
	// 创建存int的栈
	intStack := NewStack[int]()

	// 入栈
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	intStack.Print()

	// 查看栈顶
	if val, ok := intStack.Peek(); ok {
		fmt.Printf("Stack top element: %d\n", val)
	}

	// 出栈
	val, _ := intStack.Pop()
	fmt.Printf("Stack just popped: %d\n", val)
	intStack.Print()

	// 测试字符串栈
	fmt.Println("\n--- 测试字符串栈(String Stack) ---")
	// 创建存储string的栈
	strStack := NewStack[string]()
	strStack.Push("Hello")
	strStack.Push("World")
	strStack.Push("Golang!")

	// 循环出栈
	for !strStack.IsEmpty() {
		v, _ := strStack.Pop()
		fmt.Printf("Popped: %s\n", v)
	}
}
