package main

import (
	"fmt"
)

func main() {
	var i interface{}
	fmt.Println(i == nil) // true，因为 type 和 value 都为 nil

	var b *byte
	fmt.Println(b == nil) // true，指针没有分配指向的对象

	i = b
	fmt.Println(i == nil) // false, 因为 interface 的 type 为 *byte，不为 nil

	var i1 interface{}
	i = i1
	fmt.Println(i == nil) // true，因为 type 和 value 都为 nil
}

/*
	interface 实际上不是我们认为的只有简单的值，interface 分为两种类型： runtime.eface 和 runtime.iface
	eface 就是没有任何方法的空接口，是 empty interface 的简称
	iface 是包含方法的接口

	interface 的结构体如下：
		type eface struct {
			_type *_type
			data  unsafe.Pointer
		}

		type iface struct {
			tab  *itab
			data unsafe.Pointer
		}
	分别存在两个指针，一个指向数据类型，一个指向数据，因为它是 interface{} 类型的，所以需要记录对应数据的类型，这里就是使用这个类型指针 _type 来记录
	我们可以理解 interface 的类型推断就是利用这个 _type 来实现的
	对于 interface 的 nil 判断，只有 type 和 value 同时为 nil 才会为 true
*/
