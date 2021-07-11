package main

import (
	"fmt"
	"unsafe"
)

type People1 struct {
	age     int32
	address string
}
type T struct {
	a bool  //1
	e byte  //1
	c int8  //1
	b int32 //4
	d int64 //8
	// 15
	f string
	g []int64
	h People1
	i rune
	j uintptr
}

func main() {
	var t T

	fmt.Printf("t 中实际占用内存大小：%d， 对齐保证：%d\n", unsafe.Sizeof(t), unsafe.Alignof(t))
	fmt.Printf("a 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.a), unsafe.Alignof(t.a), unsafe.Offsetof(t.a))
	fmt.Printf("e 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.e), unsafe.Alignof(t.e), unsafe.Offsetof(t.e))
	fmt.Printf("c 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.c), unsafe.Alignof(t.c), unsafe.Offsetof(t.c))
	fmt.Printf("b 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.b), unsafe.Alignof(t.b), unsafe.Offsetof(t.b))
	fmt.Printf("d 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.d), unsafe.Alignof(t.d), unsafe.Offsetof(t.d))
	fmt.Printf("f 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.f), unsafe.Alignof(t.f), unsafe.Offsetof(t.f))
	fmt.Printf("g 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.g), unsafe.Alignof(t.g), unsafe.Offsetof(t.g))
	fmt.Printf("h 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.h), unsafe.Alignof(t.h), unsafe.Offsetof(t.h))
	fmt.Printf("i 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.i), unsafe.Alignof(t.i), unsafe.Offsetof(t.i))
	fmt.Printf("j 占用内存大小：%d， 对齐保证：%d, 内存地址偏移量：%d\n", unsafe.Sizeof(t.j), unsafe.Alignof(t.j), unsafe.Offsetof(t.j))

}
