package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	/*
	   dp
	   dp[i][j] 表示 [0, i] 能得到 j 的方法数
	   由于我们并不知道 j 的上下限，因此我们需要使用 map 代理 [][]int
	   由于我们并不知道上一个值的上下限，因此我们只能扫描 i-1 的 所有值
	*/
	l := len(nums)
	m := map[int]map[int]int{}

	for i, val := range nums {
		m[i] = map[int]int{}
		curMap := m[i]

		if i == 0 {
			if val == 0 {
				curMap[val] = 2
			} else {
				curMap[val] = 1
				curMap[-val] = 1
			}
			continue
		}

		//扫描上一个结果
		for k, v := range m[i-1] {
			helper(k+val, v, curMap)
			helper(k-val, v, curMap)
		}
	}

	func() {

	}()

	return m[l-1][target]
}

func helper(curVal int, preTime int, m map[int]int) {
	// 获取 i 中已经存在的 curVal 的出现次数
	time := m[curVal]
	// 出现次数+1
	time += preTime
	m[curVal] = time
}

func main() {
	sl := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	target := 1
	fmt.Println(findTargetSumWays(sl, target))
}
