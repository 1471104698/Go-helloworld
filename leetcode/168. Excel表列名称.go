package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(701))
}

func convertToTitle(columnNumber int) string {
	/*
	   columnNumber % 26 得到末尾的数字
	   columnNumber / 26 % 26
	   ZY = 26*26 + 25
	   AA = 26*1 + 1
	   BA = 26*2 + 1
	   CBA = 26*26*3 + 26*2 + 1
	*/
	bs := []byte{}
	for columnNumber > 0 {
		b := columnNumber % 26
		var ch byte
		if b == 0 {
			ch = 'Z'
		} else {
			ch = byte(b + 'A' - 1)
		}
		bs = append([]byte{ch}, bs...)
		if b == 0 {
			columnNumber -= 26
		} else {
			columnNumber -= b
		}
		columnNumber /= 26
	}
	return string(bs)
}
