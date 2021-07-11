package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type P struct {
	a int
	B int
	g G
}

type G struct {
	c int
}

// v Kind 为 ptr， ve Kind 为 struct，即调用过 Elem()
func test1(t reflect.Type, v, ve reflect.Value) {
	var vPointer uintptr
	// 通过 t 获取变量名，偏移量、在结构体中的索引位置
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		// 获取变量信息
		//fieldType := t.Field(i)
		// 调用 Field() 要求 Kind() 必须是 struct
		fieldValue := ve.Field(i)
		// 判断是否是私有变量
		if !fieldValue.CanSet() {
			// 私有变量，获取结构体地址
			if vPointer == 0 {
				vPointer = v.Pointer()
			}
			// 获取变量所在的地址
			//fPointer := vPointer + fieldType.Offset
			//// 转换为 unsafe.Pointer
			//unsafePointer := (fPointer)
		}
	}
}

func main() {
	// 测试反射修改私有变量
	// 一般情况下，反射只能调用私有变量，无法修改私有变量（反射无法调用私有方法），不过我们这里可以借助 unsafe.Pointer 来修改
	p := P{}
	v := reflect.ValueOf(&p)
	ve := v.Elem()
	a := ve.FieldByName("a")
	b := ve.FieldByName("B")
	fmt.Println(a.CanSet()) // false
	fmt.Println(b.CanSet()) // true

	val := v.Interface()
	fmt.Println(val)
	fmt.Printf("%T\n", val)

	ptype := reflect.TypeOf(p)
	atype, _ := ptype.FieldByName("a")
	fmt.Println(atype.Type)
	fmt.Printf("变量名：%v\n", atype.Name)
	test1(ptype, v, ve)

	// Pointer() 返回的是指针指向的值的地址， v 必须是指针、map、slice、func、unsafe.Pointer，不能是基本数据类型或者结构体
	fmt.Println(v.Pointer())
	//fmt.Println(ve.Pointer()) // 报错
	fmt.Printf("%v\n", uintptr(unsafe.Pointer(&p)))

	var i int
	var ip *int = &i
	it := reflect.ValueOf(ip)
	fmt.Println(it.Elem().Addr())
	fmt.Println(ip)
	fmt.Println(&i)

	v.Pointer()

	sl := []int{1, 2}
	slv := reflect.ValueOf(&sl)
	//slv = reflect.Append(slv, reflect.ValueOf(1))
	//fmt.Println(slv)
	//fmt.Println(sl)
	slv.Elem().SetLen(1)
	fmt.Println(slv)
	fmt.Println(sl)

	m := map[int]string{
		99: "stringaaaaa",
		2:  "age",
	}
	mv := reflect.ValueOf(m)
	mvk := mv.MapIndex(reflect.ValueOf(99))
	fmt.Println(mvk)
	mvks := mv.MapKeys()
	for _, v := range mvks {
		fmt.Println("v：", v)
	}
	mv.SetMapIndex(reflect.ValueOf(2), reflect.ValueOf("test"))
	fmt.Println(m)
}
