package main

import (
	"fmt"
)

/*
	函数内引用了函数外的变量，比如父函数 reutrn 一个子函数，而在子函数内部引用了父函数的变量
*/

func app() func(string) func(string) {
	t := "Hi"
	c := func(b string) func(string) {
		t = t + " " + b
		fmt.Println("第二层：", t)
		return func(c string) {
			t = t + " " + c
			fmt.Println("第三层：", t)
		}
	}
	return c
}

func main() {
	a := app()
	b := a("a")
	b("c")
	a("a1")
}
