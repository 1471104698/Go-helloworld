package main

import (
	"fmt"
	"sort"
)

func largestNumber(cost []int, target int) string {
	/*
	   要求最终的数值越大，那么就是长度要越长
	   能够添加的数字越多越好，对于能够相同长度的，那么买最大的几个数字，然后降序排序，即可得到结果

	   因此第一步先算出最多能够买多少个数字
	   有点像 完全背包问题，target 为容量，cost 为物品，求刚好塞满背包最多能放多少件物品

	   实际上我们可以在计算长度的过程中，将满足条件的子字符串都给存储起来，这样在计算到最后的时候，我们就有了所有满足最大长度的几个候选字符串
	*/

	l := len(cost)
	dp := make([]int, target+1, target+1)
	dpStr := make([][]string, target+1, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = -1
		dpStr[i] = []string{}
	}

	dp[0] = 0
	dpStr[0] = append(dpStr[0], "")

	// 初始化第一件物品的情况
	bs := []byte{}
	for i, j := cost[0], 1; i <= target; i = i + cost[0] {
		dp[i] = j
		bs = append(bs, '1')
		dpStr[i] = append(dpStr[i], string(bs))
		j++
	}

	for i := 1; i < l; i++ {
		for j := cost[i]; j <= target; j++ {
			// dp[j] 目前是前 j-1 件物品的长度，dpStr[j] 是前 j-1 件物品构成的字符串候选集合
			// 如果不选当前物品的话，那么不变
			// 如果选当前物品的话，

			// 选
			if cost[i] <= j && dp[j-cost[i]] != -1 {
				if dp[j] <= dp[j-cost[i]]+1 {
					// 舍弃掉当前所有的字符串，选择 dpStr[j-cost[i]]
					dp[j] = dp[j-cost[i]] + 1
					// 清空 dpStr[j]，将 dpStr[j] 的数据转换为 dpStr[j-cost[i]] 的所有候选字符串 加上当前字符
					dpStr[j] = []string{}
					for _, v := range dpStr[j-cost[i]] {
						newStr := string(append([]byte(v), '0'+byte(i+1)))
						dpStr[j] = append(dpStr[j], newStr)
					}
				} else if dp[j] == dp[j-cost[i]]+1 {
					for _, v := range dpStr[j-cost[i]] {
						newStr := string(append([]byte(v), '0'+byte(i+1)))
						dpStr[j] = append(dpStr[j], newStr)
					}
				}
			}
		}
	}
	if dp[target] == -1 {
		return "0"
	}
	// 计算每个候选字符串各个数字出现的次数
	res := ""
	for _, v := range dpStr[target] {
		bs := []byte(v)
		sort.Sort(sortBytes(bs))
		str := string(bs)
		if str > res {
			res = str
		}
	}
	return res
}

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

func main() {
	cost := []int{7, 6, 5, 5, 5, 6, 8, 7, 8}
	target := 12
	fmt.Println(largestNumber(cost, target))
}
