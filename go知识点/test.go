package main

import (
	"fmt"
)

func main() {
	/*
		a b
		c d
	*/
	fmt.Println(colorTheGrid(2, 2))
}
func colorTheGrid(m int, n int) int {
	/*
	   dp
	   每个格子只关心右边和上边的格子
	*/
	dp := make([][][]int, m+1, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([][]int, n+1, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = make([]int, 3, 3)
		}
	}

	mod := int(1e9) + 7
	dp[1][1][0] = 1
	dp[1][1][1] = 1
	dp[1][1][2] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			} else if i == 1 {
				// 看右边
				dp[i][j][0] = dp[i][j-1][1] + dp[i][j-1][2]
				dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][2]
				dp[i][j][2] = dp[i][j-1][0] + dp[i][j-1][1]
			} else if j == 1 {
				// 看上边
				dp[i][j][0] = dp[i-1][j][1] + dp[i-1][j][2]
				dp[i][j][1] = dp[i-1][j][0] + dp[i-1][j][2]
				dp[i][j][2] = dp[i-1][j][0] + dp[i-1][j][1]
			} else {
				// 看右边和上边
				dp[i][j][0] = (dp[i][j-1][1] + dp[i][j-1][2]) * (dp[i-1][j][1] + dp[i-1][j][2] - dp[i-1][j-1][1] - dp[i-1][j-1][2])
				dp[i][j][1] = (dp[i][j-1][0] + dp[i][j-1][2]) * (dp[i-1][j][0] + dp[i-1][j][2] - dp[i-1][j-1][0] - dp[i-1][j-1][2])
				dp[i][j][2] = (dp[i][j-1][0] + dp[i][j-1][1]) * (dp[i-1][j][0] + dp[i-1][j][1] - dp[i-1][j-1][0] - dp[i-1][j-1][1])
			}
			for k := 0; k < 3; k++ {
				dp[i][j][k] %= mod
			}
		}
	}
	return ((dp[m][n][0]+dp[m][n][1])%mod + dp[m][n][2]) % mod
}
