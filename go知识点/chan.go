package main

import (
	"fmt"
	"time"
)

func main() {
	test()
	time.Sleep(time.Hour)
}

func test() {
	ch := make(chan func())
	ch2 := make(chan func())
	go func() {
		for k := 0; k < 2; k++ {
			select {
			case i := <-ch:
				fmt.Println(i)
			case j := <-ch2:
				fmt.Println(j)
			case <-time.After(time.Duration(0)):
				fmt.Println("cao")
			}
		}
	}()
}

func test2() {
	go func() {
		time.Sleep(time.Hour)
	}()
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
}
