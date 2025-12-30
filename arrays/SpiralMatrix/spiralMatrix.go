package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 { // 处理空矩阵的情况
		return []int{}
	}

	// 初始化四个边界
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	// 准备结果数组
	res := make([]int, 0, len(matrix)*len(matrix[0]))

	for top <= bottom && left <= right {
		// 先向右
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++ // 上边界下移

		// 再向下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top > bottom || left > right { // 边界交叉检查
			break
		}

		// 再向左
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom-- // 下边界上移

		// 最后向上
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++ // 左边界右移
	}

	return res
}

func main() {

}
