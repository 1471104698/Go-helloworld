package main

import (
	"container/heap"
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	param_2 := obj.FindMedian()
	fmt.Println(param_2)
	obj.AddNum(3)
	param_2 = obj.FindMedian()
	fmt.Println(param_2)
}

type MedianFinder struct {
	MaxHeap *MaxHeap
	MinHeap *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{}
	mf.MaxHeap = &MaxHeap{}
	mf.MinHeap = &MinHeap{}
	// 不存在初始值，不需要调用 heap.Init()
	return mf
}

/*
	大顶堆存放小的那部分，小顶堆存放大的那部分
	大顶堆堆顶元素是小部分的最大值，小顶堆堆顶元素是大部分的最小值，
	大顶堆堆顶元素 <= 小顶堆堆顶元素
	大顶堆元素个数 == 小顶堆元素个数 || 大顶堆元素个数 == 小顶堆元素个数 + 1

	每次存放先存放到大顶堆，然后再将堆顶元素放到小顶堆，，如果小顶堆元素超过大顶堆元素个数+1，那么将再将小顶堆的堆顶元素放到大顶堆
		max 	min
3				3
1		1		3
2		1		2 3
4		1 2 	3 4
*/
func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.MinHeap, num)
	if this.MinHeap.Len() >= this.MaxHeap.Len() {
		heap.Push(this.MaxHeap, heap.Pop(this.MinHeap))
	}
	if this.MaxHeap.Len() >= this.MinHeap.Len()+1 {
		heap.Push(this.MinHeap, heap.Pop(this.MaxHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.MinHeap.Len() == this.MaxHeap.Len() {
		return float64(this.MinHeap.Peek()+this.MaxHeap.Peek()) / 2.0
	} else {
		return float64(this.MinHeap.Peek())
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
