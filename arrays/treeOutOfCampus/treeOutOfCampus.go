package main

import (
	"fmt"
)

// 洛谷：校门外的树，go题解
func main() {
	var l, m int
	fmt.Scan(&l, &m)

	isRemoved := make([]bool, l+1) // bool数组默认值都为false

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		for j := u; j <= v; j++ {
			isRemoved[j] = true
		}
	}

	remain := 0

	for i := 0; i <= l; i++ {
		if !isRemoved[i] {
			remain++
		}
	}

	fmt.Println(remain)
}
