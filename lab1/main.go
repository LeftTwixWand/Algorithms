package main

import (
	"fmt"

	"./binaryheap"
	"./linkedlist"
)

func useLinkedList() {

	myList := linkedlist.LinkedList{}

	myList.PushBackDouble(2)
	item := myList.PushBackDouble(3)
	myList.PushFront(1)
	myList.InsertAt(item, 4)
	myList.Print()
	fmt.Println("Count items = ", myList.Count)

	myList.DeleteAt(item.Next)
	myList.Print()
	fmt.Println("Count items = ", myList.Count)
}

func useBinaryHeap() {

	myHeap := binaryheap.BinaryHeap{}

	myHeap.Push(2)
	myHeap.Push(2)
	myHeap.Push(3)
	myHeap.Push(3)
	myHeap.Push(1)

	myHeap.Print()
}

func main() {

}
