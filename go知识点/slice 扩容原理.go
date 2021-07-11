package main

import "fmt"

func main() {
	s := []int32{1, 2}
	s = append(s, 3, 4, 5)
	fmt.Printf("len=%d, cap=%d", len(s), cap(s))

}
