package main

import "math"

func maxProfit(prices []int) int {
	/*
		只有一次买入卖出的机会
		have[i] 表示持有股票时的最大利润，由于只有一次机会，没有前置利润，因此此时它的利润必定为负
		no[i] 表示目前手中没有股票，即已经卖出股票时的最大利润，利润为 卖出股票价钱-买入股票钱

		have[i] 表示第 i 天它是持有股票的，第 i 天持有股票有两种情况：
		1、第 i 天买入的，那么利润为 -price[i]
		2、[0,i-1] 之间任意一天买入的，那么利润为 have[i-1]

		no[i] 表示第 i 天是没有持有股票的，那么第 i 没有股票的情况有两种：
		1、第 i 天卖出的，那么利润为 price[i] - have[i-1]
		2、[1, i-1] 天内卖出的，那么利润为 no[i-1]

		需要说明的是，我们 no[i] 设定的是卖出后的情况，那么就意味着它必定会经过交易，因此可能存在经过交易后利润为负数的情况，因此这种情况是不进行股票买卖
		那么我们需要最终在 max(0, no[len-1]) 取最大值

		状态压缩：
		我们可以看到，状态变量只能 i-1 天相关，因此我们只需要保存 i-1 天的状态即可，不使用数组
	*/

	l := len(prices)
	if l < 2 {
		return 0
	}
	have := make([]int, l, l)
	no := make([]int, l, l)

	have[0] = -prices[0]
	for i := 1; i < l; i++ {
		// 这里 have[i-1] 和 -prices[i] 取最大值是因为对于 have[i] 来说今天才持有股票，那么肯定是在 [0, i] 之间最低价时候买入
		have[i] = int(math.Max(float64(have[i-1]), float64(-prices[i])))
		no[i] = int(math.Max(float64(no[i-1]), float64(prices[i]+have[i-1])))
	}
	if no[l-1] < 0 {
		return 0
	}
	// 上面已经默认 no[0] 的初始值为 0，因此自动是将 0 进行比较过了，不会出现负数
	return no[l-1]
}

func main() {

}
