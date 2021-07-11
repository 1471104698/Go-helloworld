package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			time.Sleep(1)
			fmt.Println(i)
		}()
	}
	var j = 0
	for i := 0; i < 10000000; i++ {
		j += i * i * i
	}
	var ch = make(chan int)
	<-ch
}

func test_GOMAXPROCS() {
	// runtime.GOMAXPROCS() 设置目前用来并行计算的 CPU 核数，并返回上一次 CPU 核数的值，如果设置为 0 的话，那么仍然沿用当前 CPU 核数值
	// 这里设置 CPU 核数为 1，但返回的是 12
	threads := runtime.GOMAXPROCS(1)
	fmt.Println(threads)
	// 这里设置 CPU 核数为 3，但返回的是 1
	threads = runtime.GOMAXPROCS(3)
	fmt.Println(threads)
}
