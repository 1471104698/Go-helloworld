package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	/*
	   考虑业务逻辑

	   1、前后空格都去掉
	   2、除去空格后的第一个字符必须是数字或者正负号
	   3、第二个字符往后都必须是数字，遇到数字外其他字符则停止
	   4、中间过程中结果可能会超出 int(32bit)，因此需要进行判断，如果是正数界限为 math.MaxInt32，如果是负数界限为 math.MinInt32
	*/

	// 去除前后空格
	s = strings.Trim(s, " ")
	if s == "" {
		return 0
	}
	// 判断第一个字符
	i := 0
	if !isNum(s[0]) && !isOp(s[0]) {
		return 0
	}
	// 获取操作符
	op := byte('+')
	if isOp(s[0]) {
		op = s[0]
		i++
	}

	max := math.MaxInt32
	min := math.MinInt32
	res := 0
	//
	for ; i < len(s); i++ {
		if !isNum(s[i]) {
			break
		}
		val := int(s[i] - '0')
		if op == '+' && (res > max/10 || res == max/10 && val > 7) {
			return max
		}
		if op == '-' && (res > max/10 || res == max/10 && val > 8) {
			return min
		}
		res = res*10 + val
	}
	if op == '-' {
		return -res
	}
	return res
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func isOp(b byte) bool {
	return b == '+' || b == '-'
}

func main() {
	fmt.Println(myAtoi("-91283472332"))
}
