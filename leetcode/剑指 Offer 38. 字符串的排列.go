package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {
	/*
	   dfs
	   内部可能存在重复的元素，因此需要涉及到去重
	*/
	res = make([]string, 1)
	bs := []byte(s)
	sort.Sort(Bs(bs))
	dfs(bs, &[]byte{}, make([]bool, len(bs), len(bs)), 0)
	return res
}

var res []string

func dfs(bs []byte, curBs *[]byte, used []bool, k int) {
	l := len(bs)
	if k == l {
		res = append(res, string(*curBs))
		return
	}
	// 当前物色第 k 个
	for i := 0; i < l; i++ {
		// 判断是否需要跳过这一个
		if used[i] || i > 0 && bs[i] == bs[i-1] && !used[i-1] {
			continue
		}
		*curBs = append(*curBs, bs[i])
		used[i] = true
		dfs(bs, curBs, used, k+1)
		*curBs = (*curBs)[:k]
		used[i] = false
	}
}

type Bs []byte

// Len 获取切片长度
func (a Bs) Len() int {
	return len(a)
}

// Swap 交换切片元素
func (a Bs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less 元素比较逻辑
func (a Bs) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	permutation("aab")
	fmt.Println(res)
}
