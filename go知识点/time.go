package main

import (
	"fmt"
	"time"
)

// 时间日期格式
const (
	StringFormat        = "2006-01-02 15:04:05"
	StringFormatWithDay = "2006-01-02"
)

func main() {
	// 时间戳
	timeStamp := time.Now().Unix()
	fmt.Println(timeStamp)
	// 时间戳转 time.Time
	t := time.Unix(timeStamp, 0)
	fmt.Println(t)
	// time.Time 转 string
	tStr := t.Format(StringFormat)
	fmt.Println(tStr)
	// string 转 time.Time
	t, _ = time.Parse(StringFormat, tStr)
	fmt.Println(t)
}
