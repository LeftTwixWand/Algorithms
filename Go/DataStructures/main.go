package main

import (
	"fmt"
	"time"

	"github.com/LeftTwixWand/Algorithms/Go/DataStructures/binaryheap"
	"github.com/LeftTwixWand/Algorithms/Go/DataStructures/binarysearchtree"
	"github.com/LeftTwixWand/Algorithms/Go/DataStructures/binarytree"
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
	dictionary.Add(hashtable.NewKeyValuePair(0, 0))
	dictionary.Add(hashtable.NewKeyValuePair(15, 15))

	dictionary.Remove(0)

	dictionary.Print()
}

func useBinaryTree() {

	node := binarytree.NewNode(2)
	tree := binarytree.New(node)

	tree.Insert(0)

	tree.Insert(2)

	tree.Insert(4)

	tree.Insert(6)

	tree.Insert(8)
	tree.Insert(9)

	tree.Print()
}

func main() {

	// useLinkedList()
	// useBinaryHeap()
	// useHashTable()
	// useBinaryTree()

	binarysearchtree.Use()
}

// func buildPrefix(level int) string {

// 	// A format has the next structure:
// 	// lvl2   |   +---k---+
// 	// lvl1   | +-k-+   +-k-+
// 	// lvl0   |+k+ +k+ +k+ +k+
// 	// lvl -1 |k k k k k k k k

// 	// So, prefix has the next structure: lvl * 2 + 1

// 	result := ""

// 	indent := calculateFullPrefixIndent(level)
// 	centralIndex := GetCentralIndex(indent)

// 	result += GetPrefixWhiteSpaces(centralIndex)

// 	if indent > 0 {
// 		result += "+"
// 	}

// 	result += GetPrefixRelations(centralIndex)

// 	return result
// }

// func GetCentralIndex(number int) int {
// 	return (number-1)/2 + 1
// }

// func calculateFullPrefixIndent(level int) int {

// 	if level == 0 {
// 		return 0
// 	}

// 	return calculateFullPrefixIndent(level-1)*2 + 1
// }

// func GetPrefixWhiteSpaces(centralIndex int) string {

// 	prefix := ""

// 	for i := 0; i < centralIndex-1; i++ {
// 		prefix += " "
// 	}

// 	return prefix
// }

// func GetPrefixRelations(centralIndex int) string {
// 	prefix := ""

// 	for i := 0; i < centralIndex-1; i++ {
// 		prefix += "-"
// 	}

// 	return prefix
// }
