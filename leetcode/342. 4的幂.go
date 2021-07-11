package main

func isPowerOfFour(n int) bool {
	/*
	   2 的幂， n > 0 && n&(n-1) == 0

	   4：100
	   16：10000
	   1 在奇数位
	*/
	if n < 0 || n&(n-1) != 0 {
		return false
	}
	// 找到 1 的位置
	for i := 0; i < 32; i++ {
		// 判断当前位是否是 1
		if (n>>i)&1 != 0 {
			if i%2 == 0 {
				return n != 2
			}
			return false
		}
	}
	return false
}

func main() {
	isPowerOfFour(64)
}
