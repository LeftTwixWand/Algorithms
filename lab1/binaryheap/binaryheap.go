package binaryheap

import (
	"fmt"
	"math"
	"strings"

	"../enterprice"
)

// BinaryHeap is a binary search tree with a prioritet
type BinaryHeap struct {
	data []enterprice.Enterprice
}

// NewBinaryHeap is a constructor, which returns pointer on a new heap object
func NewBinaryHeap(values ...enterprice.Enterprice) *BinaryHeap {
	heap := &BinaryHeap{
		data: append([]enterprice.Enterprice{}, values...),
	}

	for i := len(heap.data)/2 - 1; i >= 0; i-- {
		heap.down(i)
	}

	return heap
}

// Count is a method which returns count of items in BinaryHeap
func (heap *BinaryHeap) Count() int {
	return len(heap.data)
}

func (heap *BinaryHeap) down(item int) {
	for {

		childItem := item*2 + 1 // select the smallest child to swap
		if childItem >= len(heap.data) || childItem < 0 {
			break
		}

		right := childItem + 1
		if right < len(heap.data) && heap.data[childItem].Profit >= heap.data[right].Profit {
			childItem = right
		}

		// swap
		if heap.data[childItem].Profit >= heap.data[item].Profit {
			break
		}

		heap.data[item], heap.data[childItem] = heap.data[childItem], heap.data[item]
		item = childItem
	}
}

// NodeAt is a method, which returns a pointer on node at current index
func (heap *BinaryHeap) nodeAt(index int) *Binder {
	if index < 0 || index >= len(heap.data) {
		return nil
	}

	return &Binder{heap: heap, index: index}
}

// Push is a method, which allows to add a new item to the heap
func (heap *BinaryHeap) Push(value enterprice.Enterprice) {
	heap.data = append(heap.data, value)
	heap.up(len(heap.data) - 1)
}

// Pop is a method, which returns the first item of heap
func (heap *BinaryHeap) Pop() (*enterprice.Enterprice, bool) {
	if len(heap.data) == 0 {
		return nil, false
	}

	val := heap.data[0]
	heap.data[0], heap.data[len(heap.data)-1] = heap.data[len(heap.data)-1], heap.data[0]
	heap.data = heap.data[:len(heap.data)-1]
	heap.down(0)

	return &val, true
}

func (heap *BinaryHeap) up(index int) {
	for {
		parentIndex := (index - 1) / 2
		if index == 0 || heap.data[parentIndex].Profit <= heap.data[index].Profit {
			break
		}
		heap.data[index], heap.data[parentIndex] = heap.data[parentIndex], heap.data[index]
		index = parentIndex
	}
}

// Root is a method, which retunrs first item of data array
func (heap *BinaryHeap) Root() *Binder {
	return heap.nodeAt(0)
}

// Height is a method, which retunrs count of levels in heap
func (heap *BinaryHeap) Height() int {
	return int(math.Floor(math.Log2(float64(len(heap.data)))))
}

// Print is a method, which allows to print out a heap
func (heap *BinaryHeap) Print() {
	for i := 0; i < len(heap.data); i++ {
		node := heap.nodeAt(i)

		// left-side whitespaces
		if leftChild := node.LeftChild(); leftChild != nil {
			if node.isRightChild() {
				fmt.Print(strings.Repeat("-", leftChild.calcPrintWidth()))
			} else {
				fmt.Print(strings.Repeat(" ", leftChild.calcPrintWidth()))
			}

		}

		// node value
		fmt.Printf(" %d ", node.Value().Profit)

		// right-side whitespaces
		if rightChild := node.RightChild(); rightChild != nil {
			if i == 0 || node.isRightChild() {
				fmt.Print(strings.Repeat(" ", rightChild.calcPrintWidth()))
			} else {
				fmt.Print(strings.Repeat("-", rightChild.calcPrintWidth()))
			}
		}

		if node.isRightMost() {
			if i == len(heap.data)-1 && !node.isRightChild() {
				if parent := node.Parent(); parent != nil {
					if vw := parent.calcValueWidth(); vw > 0 {
						fmt.Print(strings.Repeat("-", vw/2))
						fmt.Print("+")
					}
				}
			}
			fmt.Println()
		} else {
			if rightParent := node.findRightParent(); rightParent != nil {
				if vw := rightParent.calcValueWidth(); vw > 0 {
					if node.isRightChild() {
						fmt.Print(strings.Repeat(" ", vw))
					} else {
						fmt.Print(strings.Repeat("-", vw/2))
						fmt.Print("+")
						fmt.Print(strings.Repeat("-", (vw-1)/2))
					}
				}
			}
		}
	}
}

// Validate is a method, which validate a heap
func (heap *BinaryHeap) Validate() bool {
	if root := heap.Root(); root != nil {
		return root.validate()
	}
	return true
}

// Return top 25 the most profit companies
func (heap *BinaryHeap) GetTop25() []enterprice.Enterprice {

	if heap.Count() > 0 {
		var top25 []enterprice.Enterprice
		for i := 0; i < heap.Count()/4; i++ {
			top25 = append(top25, heap.data[i])
		}

		return top25
	}

	return nil
}

// Binder is a struct which works like a node in BinaryHeap
type Binder struct {
	heap  *BinaryHeap
	index int
}

// Value is a method, which returns a value of node
func (binder *Binder) Value() *enterprice.Enterprice {
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
