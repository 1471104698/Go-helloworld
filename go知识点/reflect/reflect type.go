package main

import (
	"fmt"
	"reflect"
)

type AA struct {
	b BB `json:"test"`
}

type BB struct {
}

func main() {
	a := AA{}
	fmt.Println(reflect.TypeOf(a).Field(0).Type == reflect.TypeOf(BB{}))

	fmt.Println(reflect.ValueOf(reflect.ValueOf(&AA{}).Interface()).Type().Elem().Kind())
}
