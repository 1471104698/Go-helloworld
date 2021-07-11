package main

import (
	"fmt"
	"time"
)

// Req
type Req struct {
	a string
	b string
	c string
	d string

	h time.Duration
	i int

	j bool
}

// 定义一个 Option ，它是一个 func，入参为 *Req，即需要修改参数的结构体
type Option func(req *Req)

// WithA 我们构建一个 func，这个 func 入参是用户想要修改的某个参数的值，然后利用 func 的闭包性质，出参为 Option 类型的 func，
// 我们在这个 Option 内部将要修改的参数赋值给对应的结构体变量上，这样调用这个 Option 函数就可以直接修改这个变量了
// 我们需要做的就是接收用户想要修改的变量的 Option，然后执行它们，继而完成对 Req 的初始化
func WithA(a string) Option {
	return func(req *Req) {
		req.a = a
	}
}

func WithB(b string) Option {
	return func(req *Req) {
		req.b = b
	}
}

func NewReq(opts ...Option) *Req {
	req := new(Req)
	// 执行所有用户传入的 Option，以此来设置 req 的值
	for _, opt := range opts {
		opt(req)
	}
	return req
}

func main() {
	// 我们执行 Req 提供的方法，获取对应的 Option 传入
	req := NewReq(WithA("a"), WithB("b"))
	fmt.Println(req) // &{a b   0 0 false}
}
