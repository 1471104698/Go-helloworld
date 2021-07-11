package main

import (
	"fmt"
	"time"
)

type TT struct {
}

func (*TT) eat1() {

}

func (TT) eat2() {

}

func testTT() {
	t := TT{}
	tp := &TT{}
	t.eat1()
	t.eat2()

	tp.eat1()
	tp.eat2()
}

// 时间日期格式
const (
	DefaultStringFormat        = "2006-01-02 15:04:05"
	DefaultStringFormatWithDay = "2006-01-02"
)

func main() {

	t, err := time.Parse(DefaultStringFormat, "2006-01-02 15:04:05")

	fmt.Println(t)
	fmt.Println(err)
	fmt.Println(t.Format(DefaultStringFormatWithDay))
}

type A interface {
	eat1()
	eat2()
}

func testA() {
	var t A = &TT{}
	tp := &TT{}
	t.eat1()
	t.eat2()

	tp.eat1()
	tp.eat2()
}
