package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type P struct {
	A int `json:"aaa"`
	B string
}

// 测试 StructField 中 Tag 属性
func main() {
	p := &P{}
	bs, _ := json.Marshal(p)
	fmt.Println(string(bs)) // 如果携带了 json 标签，那么对应的字段名会输出 json 标签， {"aaa":0,"B":""}

	t := reflect.TypeOf(p)
	a := t.Elem().Field(0)
	fmt.Println(a.Tag) // json:"aaa"
	bs2, _ := json.Marshal(a.Tag)
	fmt.Println(string(bs2)) // "json:\"aaa\""
}
