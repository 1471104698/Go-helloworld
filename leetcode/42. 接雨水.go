package main

func trap(height []int) int {
	/*
		两遍 for 循环，找到每个位置左边和右边最高的柱子高度
	*/
	l := len(height)
	if l < 3 {
		return 0
	}
	leftHeight := make([]int, l, l)
	rightHeight := make([]int, l, l)
	for i := 1; i < l; i++ {
		if leftHeight[i-1] > height[i-1] {
			leftHeight[i] = leftHeight[i-1]
		} else {
			leftHeight[i] = height[i-1]
		}
	}
	for i := l - 2; i >= 0; i-- {
		if rightHeight[i+1] > height[i+1] {
			rightHeight[i] = rightHeight[i+1]
		} else {
			rightHeight[i] = height[i+1]
		}
	}

	res := 0
	for i := 1; i < l-1; i++ {
		if leftHeight[i] > height[i] && rightHeight[i] > height[i] {
			if leftHeight[i] < rightHeight[i] {
				res += leftHeight[i] - height[i]
			} else {
				res += rightHeight[i] - height[i]
			}
		}
	}
	return res
}
