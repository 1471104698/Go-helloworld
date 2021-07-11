package main

import (
	"fmt"
	"math"
)

func lastStoneWeightII(stones []int) int {
	/*
	   如果是选最重的，那么只需要排序后每次选两个最重的即可，但是这里是任意
	   这里感觉可以用 dp，
	   dp[i][j] 表示 [i, j] 间的石头任意碰撞得到的最小重量

	   dp 状态转移情况：
	       在 [i, j] 内，我们假设内部的石头 k 是最后进行碰撞的，那么最终划分成了 dp[i][k-1] stones[k] 和 dp[k+1][j] 三者进行碰撞
	   dp 状态转移方程：
	       dp[i][j] = min(三者任意碰撞的最小值)
	   边界处理：如果最后碰撞的是 i 或者 j，那么只有 i 和 dp[i+1][j] 碰撞 或者 只有 j 和 dp[i][j-1] 碰撞

	*/
	l := len(stones)
	dp := make([][]int, l, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l, l)
	}

	// 值初始化
	max := 100000000
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			dp[i][j] = max
		}
	}

	//2, 23, 2

	// 处理只有一个石头的状况
	for i := 0; i < l; i++ {
		dp[i][i] = stones[i]
	}
	for j := 1; j < l; j++ {
		for i := j - 1; i >= 0; i-- {
			// 只有两个石头
			if j == i+1 {
				dp[i][j] = int(getAbs(float64(stones[i]), float64(stones[j])))
				continue
			}
			// 处理边界条件
			dp[i][j] = int(getMin(getAbs(float64(stones[i]), float64(dp[i+1][j])), getAbs(float64(stones[j]), float64(dp[i][j-1]))))

			//从 (i, j) 间选择一个最后碰撞的石头
			for k := i + 1; k < j; k++ {
				dp[i][j] = int(getMin(
					float64(dp[i][j]), getMinWitThree(float64(dp[i][k-1]), float64(dp[k+1][j]), float64(stones[k]))),
				)
			}
		}
	}

	return dp[0][l-1]
}

func getMinWitThree(i, j, k float64) float64 {
	return getMin(getAbs(i-j, k), getMin(getAbs(i-k, j), getAbs(j-k, i)))
}

func getAbs(i, j float64) float64 {
	return math.Abs(math.Abs(float64(i - j)))
}

func getMin(i, j float64) float64 {
	return math.Min(float64(i), float64(j))
}

func main() {
	stones := []int{10, 31, 30, 60, 53}
	fmt.Println(lastStoneWeightII(stones))
}
