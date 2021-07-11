package main

import "fmt"

/*
	Go 这种设计比价难受：结构体 A 内嵌结构体 B，它们之间并不是什么父子关系，即不是 Java 的父子类
	这也就导致了当 结构体 B 实现了 eat() 和 say()，并且在 eat() 中调用了 say()，而结构体 A 也同样实现了 say()
	如果这时候使用 A 的实例去调用 eat()，那么由于 A 没有实现 eat()，所以调用的都是 B 的 eat()
	而在 B eat() 内部调用的 say() 也仍然是 B 的 say()，因为 A 和 B 本身并什么父子关系
	所以对 B 来说它就不会判断这是借助 A 调用过来的，不会去判断是否需要去查找 A 是否有 say() 然后调用它

	这个点算是跟 Java 不同的一个点之一
*/
type L struct {
}

func (l *L) eat() {
	fmt.Println("eat L")
	l.say()
}

func (*L) say() {
	fmt.Println("say L")
}

type L1 struct {
	*L
}

func (*L1) say() {
	fmt.Println("say L1")
}

func main() {
	l1 := &L1{
		L: &L{},
	}
	l1.eat()
	//eat L
	//say L
}
