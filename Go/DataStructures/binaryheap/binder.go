package binaryheap

import (
	"fmt"
	"math"
)

// Binder is a struct which works like a node in BinaryHeap for printing a heap
type Binder struct {
	heap  *BinaryHeap
	index int
}

// Value is a method, which returns a value of node
func (binder *Binder) Value() *Enterprice {
	return &binder.heap.data[binder.index]
}

// Parent is a method, which returns a parrent of node
func (binder *Binder) Parent() *Binder {
	return binder.heap.nodeAt((binder.index+1)/2 - 1)
}

// LeftChild is a method, which return a pointer on the left child
func (binder *Binder) LeftChild() *Binder {
	return binder.heap.nodeAt((binder.index+1)*2 - 1)
}

// RightChild is a method, which returns a poiner on the right child
func (binder *Binder) RightChild() *Binder {
	return binder.heap.nodeAt((binder.index + 1) * 2)
}

// Depth is a method, witch returns in which level item located
func (binder *Binder) Depth() int {
	return int(math.Floor(math.Log2(float64(binder.index + 1))))
}

// Height is a method, which returns how many levels in the heap
func (binder *Binder) Height() int {
	height := 0
	if binder.LeftChild() != nil {
		leftHeight := binder.LeftChild().Height()
		height = leftHeight + 1
	}
	if binder.RightChild() != nil {
		rightHeight := binder.RightChild().Height()
		if rightHeight+1 > height {
			height = rightHeight + 1
		}
	}
	return height
}

func (binder *Binder) calcValueWidth() int {
	v := binder.Value()
	if v.Profit == 0 {
		return 3
	}

	return int(math.Log10(float64(v.Profit))) + 3
}

func (binder *Binder) calcPrintWidth() int {
	width := binder.calcValueWidth()

	if binder.LeftChild() != nil {
		width += binder.LeftChild().calcPrintWidth()
	}
	if binder.RightChild() != nil {
		width += binder.RightChild().calcPrintWidth()
	}

	return width
}

func (binder *Binder) isRightChild() bool {
	return binder.index > 0 && binder.index%2 == 0
}

func (binder *Binder) isRightMost() bool {
	return binder.index == len(binder.heap.data)-1 ||
		binder.index == int(math.Exp2(float64(binder.Depth()+1)))-2
}

func (binder *Binder) findRightParent() *Binder {
	node := binder
	for {
		if !node.isRightChild() {
			return node.Parent()
		}
		node = node.Parent()
	}
	// return nil
}

func (binder *Binder) validate() bool {
	if leftChild := binder.LeftChild(); leftChild != nil {
		if leftChild.Value().Profit < binder.Value().Profit || !leftChild.validate() {
			fmt.Printf("SELF: %d / LEFT: %d\n", binder.Value().Profit, leftChild.Value().Profit)
			return false
		}
	}
	if rightChild := binder.RightChild(); rightChild != nil {
		if rightChild.Value().Profit < binder.Value().Profit || !rightChild.validate() {
			fmt.Printf("SELF: %d / RIGHT: %d\n", binder.Value().Profit, rightChild.Value().Profit)
			return false
		}
	}
	return true
}
