package main

import (
	"fmt"
	"reflect"
)

type AAA struct {
	i *int
}

func main() {
	i := 1
	hh(&i)
}

func hh(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Println(v.Elem().Interface())
	fmt.Println(reflect.TypeOf(v.Elem().Interface()).Kind())
}
