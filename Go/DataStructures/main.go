package main

import (
	"fmt"

	// "./binaryheap"
	// "./enterprice"
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

// func useBinaryHeap() {

// 	myHeap := binaryheap.BinaryHeap{}

// 	myHeap.Push(enterprice.NewEnterprice("name1", "address1", time.January, 2001, 1))
// 	myHeap.Push(enterprice.NewEnterprice("name2", "address2", time.February, 2002, 2))
// 	myHeap.Push(enterprice.NewEnterprice("name3", "address3", time.March, 2003, 3))
// 	myHeap.Push(enterprice.NewEnterprice("name4", "address4", time.April, 2004, 4))
// 	myHeap.Push(enterprice.NewEnterprice("name11", "address11", time.May, 2011, 1))

// 	myHeap.Print()

// 	fmt.Println()
// 	fmt.Println(myHeap.GetTop25())
// }

func main() {

	useLinkedList()

}
