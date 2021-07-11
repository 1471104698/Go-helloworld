package main

import (
	"fmt"
	"reflect"
	"strings"
)

type A struct {
	name string `di:"singleton" json:"fuck"`
}

func main() {
	t := reflect.TypeOf(A{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Tag.Get("di"))
	}
}

// isDI 判断是否需要注入
func isDI(tag string) bool {
	return strings.Contains(tag, "di")
}
