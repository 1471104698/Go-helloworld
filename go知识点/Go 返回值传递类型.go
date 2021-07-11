package main

import "fmt"

type AAA struct {
	b BBB
}

func (a AAA) getB() BBB {
	return a.b
}

type BBB struct {
	age int
}

func main() {
	a := AAA{}
	b := a.b
	b.age = 1
	fmt.Println(b)
	fmt.Println(a.b)
}
