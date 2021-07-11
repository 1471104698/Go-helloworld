package main

import "fmt"

func longestPalindrome(s string) string {
	/*
		输入：s = "babad"
		输出："bab"
		解释："aba" 同样是符合题意的答案。

		1、dp
			dp[i][j] 有以下几种意思：
				[i, j] 之间的最长回文串长度
				[i, j] 作为边界构成的回文串长度,构不成为 -1
				[i, j] 是否能够构成回文串，此时存储的是 true/false，使用额外的一个变量来记录最长的回文串长度
			第二种和第三种实际上是一个类型，第一种往后移动新添加进来的字符建立不了联系

		2、中心扩展
			以 [i,i] 和 [i,i+1] 作为中心点向左右两边扩展
	*/

	//l := len(s)
	//dp := make([][]bool, l, l)
	//for i := 0; i < l; i++ {
	//	dp[i] = make([]bool, l, l)
	//	dp[i][i] = true
	//}
	//idx := 0
	//mlen := 1
	//for j := 1; j < l; j++ {
	//	for i := j - 1; i >= 0; i-- {
	//		if s[i] == s[j] && (dp[i+1][j-1] || i+1 == j) {
	//			curLen := j - i + 1
	//			dp[i][j] = true
	//			if mlen < curLen {
	//				mlen = curLen
	//				idx = i
	//			}
	//		}
	//	}
	//}
	//return s[idx : idx+mlen]

	l := len(s)
	if l == 0 {
		return ""
	}
	leftIdx := 0
	mlen := 1
	for i := 0; i < l; i++ {
		l1 := isPart(s, i, i)
		l2 := isPart(s, i, i+1)
		if l1 > mlen {
			mlen = l1
			leftIdx = i - l1/2
		}
		if l2 > mlen {
			mlen = l2
			leftIdx = i - (l2-1)/2
		}
		// c b a b
		// 0 1 2 3 len = 3, i = 2, idx = 1
		// b a a b
		// 0 1 2 3 len = 4, i = 1, idx = 0
	}

	return s[leftIdx : leftIdx+mlen]
}

func isPart(s string, i, j int) int {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return j - i - 1
}

func main() {
	fmt.Println(longestPalindrome("aacabdkacaa"))
}
