package main

import (
	"fmt"
	"math"
)

func findMaxForm(strs []string, m int, n int) int {
	/*
		01 背包问题，strs 是物品，m n 是容量

		dp
		dp[j][a][b] 表示 前 j 个 str 中最多具有 a 个 0， b 个 1 的最大子集个数


		1、状态转移情况：
			zero[k] 表示第 k 个字符串中 0 的个数
			one[k] 表示第 k 个字符串中 1 的个数
			第 j 个 str 可选可不选

		状态转移方程：
			dp[j][a][b] = max(dp[j-1][a-zero[j]][b-one[j]]+1, dp[j-1][a][b])


		2、状态转移方程压缩情况：
			对于 01 背包问题，我们可以省略掉物品那一维度，只考虑背包容量
			这一种情况我们需要倒叙遍历背包容量。
			因为这样我们在遍历时间维度为 T， 背包容量为 S 的时候， 我们只需要时间维度为 T-1， 背包容量 < S 的情况，
			倒叙遍历的话只会覆盖掉 T-1 中背包容量 > S 的情况，因此可以重复利用

			dp[a][b] 表示选择前 i 个 str 时，最多具有 a 个 0， b 个 1 的最多子集数

		状态转移压缩方程：
			for i := 0; i < len(strs); i++ {
				dp[a][b] = max(dp[a][b], dp[a-zero][b-one]+1)
			}

	*/
	l := len(strs)
	dp := make([][]int, m+1, m+1)
	for j := 0; j <= m; j++ {
		dp[j] = make([]int, n+1, n+1)
	}

	for j := 0; j < l; j++ {
		zero, one := 0, 0
		for _, v := range strs[j] {
			if v == '0' {
				zero++
			} else {
				one++
			}
		}
		for a := m; a >= 0; a-- {
			for b := n; b >= 0; b-- {
				// 那么不存在上一个 str，跟上一个 str 无关，因此需要考虑当前 str 是否能装，不能为 0，能为 1
				if j == 0 {
					if zero <= a && one <= b {
						dp[a][b] = 1
					}
					continue
				}
				// 如果当前 str 的 0 1 个数超过背包容量
				if zero > a || one > b {
					// 跳过当前 str
					continue
				}
				dp[a][b] = int(math.Max(float64(dp[a-zero][b-one]+1), float64(dp[a][b])))
			}
		}
	}
	return dp[m][n]
}

func main() {
	strs := []string{"0", "11", "1000", "01", "0", "101", "1", "1", "1", "0", "0", "0", "0", "1", "0", "0110101", "0", "11", "01", "00", "01111", "0011", "1", "1000", "0", "11101", "1", "0", "10", "0111"}
	m, n := 9, 80
	fmt.Println(findMaxForm(strs, m, n))
}
