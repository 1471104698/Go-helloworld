package main

import (
	"fmt"
	"math"
)

func maxProfitIII(prices []int) int {
	/*
		最多完成两笔交易，那么相比一次交易来说，第二次交易是具有前置利润的
		注意第二次交易的买入需要在第一次交易的卖出之后，并且卖出当天不能再进行买入，即第一次交易的卖出和第二次交易的买入不能同一天进行

		have[i][0] 表示第 i 天或之前就已经买入了股票，并且是第一次交易，对于这一次交易来说，它是没有前置利润的
		have[i][1] 表示第 i 天或之前就已经买入了股票，并且是第二次交易，对于这一次交易来说，它已经经过了第一次交易，第一次交易的利润为 no[i-1][0]
		no[i][0] 表示第 i 天或者之前就已经卖出了股票，并且是第一次交易，卖出时利润为 prices[i] + have[i-1][0]
												(这里利润的计算为什么是 + 呢? 因为我们将 have[i][0] 买入时的利润设置为 -prices[i])
		no[i][1] 表示第 i 天或者之前就已经卖出了股票，并且在前面已经卖出过一次股票了，是第二次交易，总的利润是 prices[i] + have[i][1]

		需要注意的是，可能整个数组中只有一次交易能够赚到钱，如果强制进行第二次交易那么会亏钱，也可能两次交易都能赚钱，也可能一次交易都赚不了钱，
		因此我们结果取 max(0, no[len-1][0] no[len-1][1])

		状态压缩：
		依赖是跟 i-1 有关，可以压缩
	*/
	l := len(prices)
	if l < 2 {
		return 0
	}
	have := make([][]int, l, l)
	for i := 0; i < l; i++ {
		have[i] = make([]int, 2, 2)
	}
	no := make([][]int, l, l)
	for i := 0; i < l; i++ {
		no[i] = make([]int, 2, 2)
	}
	have[0][0] = -prices[0]
	// 这里目的是让第二次交易买入或者卖出的时候不取默认值 0，而是去取第一次交易的值
	have[1][1] = math.MinInt32
	no[1][1] = math.MinInt32
	for i := 1; i < l; i++ {
		have[i][0] = int(math.Max(float64(have[i-1][0]), float64(-prices[i])))
		no[i][0] = int(math.Max(float64(no[i-1][0]), float64(prices[i]+have[i-1][0])))
		if i >= 2 {
			have[i][1] = int(math.Max(float64(have[i-1][1]), float64(no[i-1][0]-prices[i])))
			no[i][1] = int(math.Max(float64(no[i-1][1]), float64(prices[i]+have[i-1][1])))
		}
	}

	// 同样初始值为 0，不需要判断负数情况
	if no[l-1][0] > no[l-1][1] {
		return no[l-1][0]
	} else {
		return no[l-1][1]
	}
}

func main() {
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfitIII(prices))
}
