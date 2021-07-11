package main

import "fmt"

func numSquares(n int) int {
	/*
		   贪心试试，二分法找到最接近 n 的完全平方根

		   贪心不行，比如 12，那么贪心它会找 9 + 1 + 1 + 1，但实际上最佳的结果是 4 + 4 + 4
		因此还是需要 dp，双重 for 循环 dp

		目前下面的代码是贪心的，是错误的，后面需要改成 dp
	*/
	c := 0
	for n > 0 {
		c++
		n -= h(n)
	}
	return c
}

// 找到最接近 n 的完全平方数
func h(n int) int {
	left := 0
	right := n
	for left < right {
		mid := (left + right + 1) / 2
		if mid*mid > n {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left * left
}

func main() {
	fmt.Println(numSquares(12))
}
