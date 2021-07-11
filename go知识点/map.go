package main

import "fmt"

/*
	问题①：map 是无序集合，但是一般情况下数据存储的位置是固定的，为什么每次输出的顺序都是不固定的？

	解答：因为 for range 编译器会调用一个随机函数生成一个随机值，作为 map 开始遍历的初始位置
	每次 for range 生成的随机数都不同，因此输出的结果顺序是不固定的

	理由：
	目前看到的是官方有意而为之的，因为 map 每次扩容都会导致 key-value 的位置发生改变，为了避免用户依赖顺序，因此直接打乱
*/
func main() {
	m := map[int32]int32{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	for i := 0; i < 3; i++ {
		for _, v := range m {
			fmt.Printf("%v\t", v)
		}
		fmt.Println()
	}

}
