package main

import "fmt"

func main() {
	deadends := []string{
		"0201", "0101", "0102", "1212", "2002",
	}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}

func openLock(deadends []string, target string) int {
	/*
	   bfs，假设最多走 k 步，那么我们需要从第 1 步开始记录每一步可以走的序列集合，将这个序列集合作为下一步选择的基础
	*/

	// 由于 goalng 没有 set 集合，所以这里直接将使用 map 代替，反正 Java 中 set 底层也是使用 map
	// 这里的 struct{}{} 可以理解为代替 value 的作用，用来借助判断 key 是否存在的空结构
	// 使用 map 方便查找某个 target 是否存在于 deadends 中
	deadendsSet := map[string]struct{}{}
	for _, v := range deadends {
		deadendsSet[v] = struct{}{}
	}
	// 开局直接王炸
	if contains(deadendsSet, "0000") {
		return -1
	}
	// 构建序列队列 queue，同时将 初始序列 "0000" 入队
	// golang 中没有显式的队列结构可以使用，这里直接使用 slice，入队在队尾，即在末尾索引位置，出队在队头，即在 0 号索引位置
	queue := []string{"0000"}

	// 记录已经出现过的序列，避免多个前几步序列已经构成相同的序列还入队了
	set := map[string]struct{}{}

	// 当前已经走过的步数
	step := 0
	for len(queue) > 0 {

		// 扫描上一步的所有字符串
		size := len(queue)
		for size > 0 {
			// 拿到上一步的序列 str
			var preStr string
			queue, preStr = queue[1:], queue[0]
			if preStr == target {
				return step
			}
			// string 不可变，要修改需要转换为 []byte
			bs := []byte(preStr)

			// 尝试每个位置的旋转
			for i := 0; i < 4; i++ {
				// +1，如果原来是 9 的话变成 0
				bs[i] = (preStr[i]-'0'+1)%10 + '0'
				checkAndPut(bs, &queue, deadendsSet, set)
				// -1，如果原来是 0 的话要变成 9，这里实际上是可以修改成 (preStr[i]-9) % 10，不过为了更明了就不改了，问题也不大
				bs[i] = ((preStr[i]-'0'-1)+10)%10 + '0'
				checkAndPut(bs, &queue, deadendsSet, set)
				// 复原当前位置，避免影响下一位的旋转
				bs[i] = preStr[i]
			}
			size--
		}
		step++
	}
	return -1
}

// contains 判断 m 中是否存在 target key
func contains(m map[string]struct{}, target string) bool {
	_, exist := m[target]
	return exist
}

// checkAndPut
func checkAndPut(bs []byte, queue *[]string, deadendsSet, set map[string]struct{}) {
	newStr := string(bs)
	if !contains(deadendsSet, newStr) && !contains(set, newStr) {
		// 将当前序列存储到已访问集合
		set[newStr] = struct{}{}
		// 入队
		*queue = append(*queue, newStr)
	}
}
