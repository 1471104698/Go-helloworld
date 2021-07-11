package main

import (
	"container/heap"
	"fmt"
)

/*
	golang 提供了一个 heap 包，该包是对 heap 的通用操作的封装，我们要使用的话只需要实现对应的接口即可
	类似于我们要排序，sort 内部实现了通用逻辑，我们只需要实现对应的特有方法逻辑即可

	heap 包的接口：
	type Interface interface {
		sort.Interface
		Push(x interface{}) // add x as element Len()
		Pop() interface{}   // remove and return element Len() - 1.
	}
	它里面依赖了 sort.Interface 接口，所以总的接口如下：
	type Interface interface {
		Len() int			// 获取长度
		Less(i, j int) bool	// 比较逻辑
		Swap(i, j int)		// 交换逻辑
		Push(x interface{})	// push 逻辑
		Pop() interface{}	// pop 逻辑
	}

	同时 heap 包内部有以下几个方法：
		func Init(h Interface)
		func Push(h Interface, x interface{})
		func Pop(h Interface) interface{}
		func Remove(h Interface, i int) interface{}
		func Fix(h Interface, i int)
	它们接收的参数存在 h Interface，即我们自己的结构体实现了它的 Interface 接口，然后调用 heap 包的这些方法将结构体传入
	它内部会自动调用我们结构体实现的逻辑来完成对应的操作，类似 sort.Sort(Interface{})
*/
func main() {
	maxHeap := &MaxHeap{}
	// 如果 maxHeap 已经存在数据，那么可以调用 heap.Init() 构建好初始堆，空数据的话就没必要
	heap.Init(maxHeap)

	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	for len(*maxHeap) > 0 {
		fmt.Println(maxHeap.Pop())
	}
}

// 利用 Heap 实现共用逻辑，做法值得学习，类似隐式的父类
type Heap []int

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

// Pop 数组末尾元素
// 注意这里为什么是 Pop 数组末尾元素，一般情况下顶堆 Pop() 的值应该是在堆顶，即 nums[0]，但是我们这里是借 heap.Pop() 来进行调用的
// 因此我们需要看 heap.Pop() 的逻辑，它是先将 nums[0] 和 nums[len-1] 进行交换，再调整堆结构，再调用这里的 Pop()
// 即它将堆顶元素放到了数组末尾，所以我们这里 Pop() 的应该是数组末尾的元素
func (h *Heap) Pop() (v interface{}) {
	// 这个做法值得学习
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return v
}

// Push 在数组末尾
// heap.Push() 是先调用这里的 Push()，再进行堆调整，因此直接放在数组末尾没毛病
func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *Heap) Peek() int {
	return (*h)[0]
}

type MaxHeap struct {
	Heap
}

// Less() 表示 i 是否需要排在 j 前面，这里的 i > j，因为 i 在 j 后面，所以这里判断是否需要将 i 交换到 j 前面
// 我们这里是大顶堆，那么越大的值越排在前面，那么如果排在后面的 i 满足 nums[i] > nums[j]，那么 i 需要交换到前面
func (h *MaxHeap) Less(i, j int) bool {
	return h.Heap[i] > h.Heap[j]
}

type MinHeap struct {
	Heap
}

// Less() 表示 i 是否需要排在 j 前面，这里的 i > j，因为 i 在 j 后面，所以这里判断是否需要将 i 交换到 j 前面
// 我们这里是小顶堆，那么越小的值越排在前面，那么如果排在后面的 i 满足 nums[i] < nums[j]，那么 i 需要交换到前面
func (h *MinHeap) Less(i, j int) bool {
	return h.Heap[i] < h.Heap[j]
}
