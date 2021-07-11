package main

import (
	"fmt"
	"unsafe"
)

func main() {
	test2()
}

func test1() {
	i := 1
	j := 2
	ip := &i
	jp := &j
	diffp := uintptr(unsafe.Pointer(jp)) - uintptr(unsafe.Pointer(ip))
	new_ip := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ip)) + diffp))
	*new_ip = 3
	fmt.Println(*new_ip)
	fmt.Println(j)
}

type Num struct {
	i string
	j int64
}

func test2() {
	n := Num{i: "EDDYCJY", j: 1}
	nPointer := unsafe.Pointer(&n)

	ptr := uintptr(nPointer)

	njPointer := (*int64)(unsafe.Pointer(ptr + unsafe.Offsetof(n.j)))
	*njPointer = 2

	fmt.Println(n)
}
