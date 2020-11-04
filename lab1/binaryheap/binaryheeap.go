package binaryheap

import "../binder"

// BinaryHeap is a binary search tree with a prioritet
type BinaryHeap struct {
	count int
	head  *binder.Binder
}

// Count is a method which returns count of items in BinaryHeap
func (heap *BinaryHeap) Count() int {
	return heap.count
}

func swap(node1, node2 *binder.Binder) {
	tmp := node1
	node1 = node2
	node2 = tmp
	tmp = nil
}

// Insert is a method which add a new item to the heap
func (heap *BinaryHeap) Insert(value int) {

	heap.count++
	if heap.head == nil {
		heap.head = binder.NewBinder(value)
		return
	}

	// tmp := heap.head
	// heap.head = binder.NewBinder(value)

}
