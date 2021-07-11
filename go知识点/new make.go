package main

/*
	new 和 make 的区别：
	 make 只作用雨 slice、map、chan，它返回的是数据本身，优点是能够初始化 len 和 cap
		make([]int, 2, 4)
		make(map[int]int, 2)
		make(chan int, 2)
		因为 map 和 chan 没有容量的概念，所以只需要指定 len
	 new 能够作用于任何数据类型，返回的是创建完后数据的指针，便捷，但是对于上述三种类型无法指定 len 和 cap
		new(int)
		new(string)
		new(T)
	一般情况下如果知道了所需的容量大小，那么可以使用 make() 来进行初始化，减少扩容

	new([]int) 等价于 &[]int{}
*/



