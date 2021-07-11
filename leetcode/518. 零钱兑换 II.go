package main

import "fmt"

func change_未压缩版(amount int, coins []int) int {
	/*
		   完全背包问题
		   硬币数量是不限制的

		   动态转移情况：
		   dp[i][j] 表示 前 i 件物品组成 j 元的方案数
		   对于每个物品有两种选择，选和不选

			动态转移方程：
			dp[i][j] = dp[i][j-1] + dp[i-coins[j]][j]
	*/

	l := len(coins)
	dp := make([][]int, l, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, amount+1, amount+1)
	}

	// 初始化只有 0 元的情况，无论什么硬币都能过构成 0 元
	for i := 0; i < l; i++ {
		dp[i][0] = 1
	}
	// 初始化只有一件物品的情况
	for i := coins[0]; i <= amount; i++ {
		dp[0][i] = dp[0][i-coins[0]]
	}

	for i := 1; i < l; i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j < coins[i] {
				continue
			}
			dp[i][j] += dp[i][j-coins[i]]
		}
	}
	return dp[l-1][amount]
}

func change_压缩版(amount int, coins []int) int {
	/*
		   完全背包问题
		   硬币数量是不限制的

		   动态转移情况：
		   dp[i][j] 表示 前 i 件物品组成 j 元的方案数
		   对于每个物品有两种选择，选和不选

			动态转移方程：
			dp[i][j] = dp[i][j-1] + dp[i-coins[j]][j]

			动态压缩，压缩掉物品这一栏，对于 T 时刻的数据来说， dp 的数据为 T-1 时刻的数据
	*/

	l := len(coins)
	dp := make([]int, amount+1, amount+1)

	dp[0] = 1
	// 初始化只有一件物品的情况
	for i := coins[0]; i <= amount; i++ {
		dp[i] = dp[i-coins[0]]
	}

	for i := 1; i < l; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	fmt.Println(change_压缩版(amount, coins))
}
