package main

import (
	"fmt"
	"reflect"
)

/*
	一个类型转换的思路：
	通过 reflect.ValueOf().Kind() 来获取它的类型信息，然后存储到 map 之类的数据结构，后面使用 switch 来进行判断
	我们可以使用 map[reflect.Kind]func() 来存储，对每个 type 绑定一个 func()，这样在 case 判断为 true 调用该函数进行某些处理
*/
type People struct {
	a int
	b int64
	C int
}

func (People) Eat() int {
	fmt.Println("eat")
	return 1
}

func (*People) Say(i int) {
	fmt.Println("say：", i)
}

func (*People) say(i int) {
	fmt.Println("say：", i)
}

func main() {
	// 1、
	//t := reflect.String
	//fmt.Printf("%T\n", t)
	//fmt.Printf("%v\n", t)
	//pt := reflect.ValueOf(People{}).Kind()
	//var is = []interface{}{1, "2", People{}}
	//for _, i := range is {
	//	switch v := reflect.ValueOf(i); v.Kind() {
	//	case t:
	//		fmt.Println("string")
	//	case pt:
	//		fmt.Println("People")
	//	default:
	//		fmt.Println("default")
	//	}
	//}

	// 2、
	//v := reflect.ValueOf(&People{})
	//v.MethodByName("Eat").Call(nil)
	//params := []reflect.Value{
	//	reflect.ValueOf(20),
	//}
	//v.MethodByName("Say").Call(params)
	//v.MethodByName("say")
	//
	//// 传入的是值
	//v1 := reflect.ValueOf(People{})
	//v1.MethodByName("Eat").Call(nil)    // eat
	//v1.MethodByName("Say").Call(params) //  panic
	//
	//v2 := reflect.ValueOf(1)
	//fmt.Println(v2.MethodByName("1")) // 空 Method{}

	// 3、
	//i := 0
	//v := reflect.ValueOf(i)
	//// 判断是否是零值，比如 int 是否是 0， string 是否是 ""，结构体判断内部变量字段是否都是零值，map、slice、chan、ptr、func、unsafe.Pointer 判断是否是 nil
	//// 需要注意的是，var is []int 这种定义方式如果直接判断的话 is == nil 为 true，但是实际上它已经创建好了 slice 结构，
	//// 不过内部指向数组的 pointer 为 nil，所以返回 nil，但是结构是存在的，可以直接 append
	//fmt.Println(v.IsZero())
	//
	//p := People{}
	//v2 := reflect.ValueOf(p)
	//fmt.Println(v2.IsZero())
	//
	//p = People{
	//	a: 1,
	//}
	//v2 = reflect.ValueOf(p)
	//fmt.Println(v2.IsZero()) // false
	//
	//var is []int
	//v3 := reflect.ValueOf(is)
	//fmt.Println(v3.IsZero()) // true
	//
	//is = append(is, 1)
	//v3 = reflect.ValueOf(is)
	//fmt.Println(v3.IsZero()) // false
	//
	//var m map[int]struct{}
	//v4 := reflect.ValueOf(m)
	//fmt.Println(v4.IsZero()) // true
	//
	//m = map[int]struct{}{}
	//v4 = reflect.ValueOf(m)
	//fmt.Println(v4.IsZero()) // true

	//// 4、
	//i := 0
	//v := reflect.ValueOf(i)
	//fmt.Println(v.IsValid()) // true
	//
	//v1 := reflect.ValueOf(nil)
	//fmt.Println(v1.IsValid()) // false
	//
	//var ip *int = nil
	//v2 := reflect.ValueOf(ip)
	//fmt.Println(v2.IsValid())        // false
	//fmt.Println(v2.Elem().IsValid()) // false
	//
	//var p People
	//v3 := reflect.ValueOf(p)
	//fmt.Println(v3.IsValid()) // false
	//
	//// 获取 p 中存在的字段 Value
	//v5 := v3.FieldByName("a")
	//fmt.Println(v5.IsValid()) // true
	//
	//// 获取 p 中不存在的字段 Value
	//v5 = v3.FieldByName("f")
	//fmt.Println(v5.IsValid()) // false

	p := People{}
	v := reflect.ValueOf(&p)
	v.Elem().FieldByName("C").Set(reflect.ValueOf(1))
	fmt.Println(p)
	v.CanSet()
}
