package main

import (
	"fmt"
	"time"

	"github.com/LeftTwixWand/Algorithms/Go/DataStructures/binaryheap"
	"github.com/LeftTwixWand/Algorithms/Go/DataStructures/hashtable"
	"github.com/LeftTwixWand/Algorithms/Go/DataStructures/linkedlist"
)

func useLinkedList() {

	fmt.Println("LinkedList:")

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

	fmt.Println()
	fmt.Println()
}

func useBinaryHeap() {

	fmt.Println("BinaryHeap:")

	myHeap := binaryheap.BinaryHeap{}

	myHeap.Push(binaryheap.NewEnterprice("name1", "address1", time.January, 2001, 1))
	myHeap.Push(binaryheap.NewEnterprice("name2", "address2", time.February, 2002, 2))
	myHeap.Push(binaryheap.NewEnterprice("name3", "address3", time.March, 2003, 3))
	myHeap.Push(binaryheap.NewEnterprice("name4", "address4", time.April, 2004, 4))
	myHeap.Push(binaryheap.NewEnterprice("name11", "address11", time.May, 2011, 1))

	myHeap.Print()

	fmt.Println()
	fmt.Println(myHeap.GetTop25())

	fmt.Println()
	fmt.Println()
}

func useHashTable() {

	dictionary := hashtable.New(15)
	result, isAdded := dictionary.Add(hashtable.NewKeyValuePair(0, 4))
	dictionary.Add(hashtable.NewKeyValuePair(1, 3))
	dictionary.Add(hashtable.NewKeyValuePair(2, 123))
	dictionary.Add(hashtable.NewKeyValuePair(3, 123124))

	fmt.Println(result)
	fmt.Println(isAdded)
	dictionary.Print()
}

func main() {

	// useLinkedList()
	// useBinaryHeap()
	useHashTable()
}
