package linkedlist

import (
	"fmt"
)

// LinkedList data structure type
type LinkedList struct {
	Count int
	first *Node
	last  *Node
}

// PushFront method
func (list *LinkedList) PushFront(item int) *Node {

	tempNode := NewNode(item)
	list.Count++

	if list.first == nil {
		list.first = tempNode
		list.last = tempNode
		return tempNode
	}

	tempNode.Next = list.first
	list.first.Previous = tempNode
	list.first = tempNode

	return tempNode
}

// PushBackSingle method implementation for single linked list
func (list *LinkedList) PushBackSingle(item int) *Node {
	tmpNode := NewNode(item)

	list.Count++
	if list.first == nil {
		list.first = tmpNode
		list.last = tmpNode
		return tmpNode
	}

	currentNode := list.first

	for ; currentNode.Next != nil; currentNode = currentNode.Next {
	}
	currentNode.Next = tmpNode
	return tmpNode
}

// PushBackDouble allows us to push new item at the back of a list
func (list *LinkedList) PushBackDouble(item int) *Node {

	tmpNode := NewNode(item)
	list.Count++

	if list.last == nil {
		list.first = tmpNode
		list.last = tmpNode
		return tmpNode
	}

	list.last.Next = tmpNode
	tmpNode.Previous = list.last
	list.last = tmpNode
	return tmpNode
}

// Print out all items
func (list *LinkedList) Print() {
	for tmpNode := list.first; tmpNode != nil; tmpNode = tmpNode.Next {
		fmt.Print(tmpNode.Value, "  ")
	}
	fmt.Println()
}

// PrintReverse out all items in reverse order
func (list *LinkedList) PrintReverse() {
	for tmpNode := list.last; tmpNode != nil; tmpNode = tmpNode.Previous {
		fmt.Print(tmpNode.Value, "  ")
	}
	fmt.Println()
}

// InsertAt is a method, which allows to insert some item after specific node
func (list *LinkedList) InsertAt(myNode *Node, item int) *Node {
	tmpNode := NewNode(item)
	tmpNode.Previous = myNode
	tmpNode.Next = myNode.Next
	myNode.Next = tmpNode
	list.Count++
	return tmpNode
}

// DeleteAt is a method, which allows to delete some item from the list
func (list *LinkedList) DeleteAt(item *Node) {

	list.Count--
	if item.Previous == nil { // If item is last
		item.Next.Previous = nil
		item = nil
		return
	}

	if item.Next == nil { // If item is first
		item.Previous.Next = nil
		item = nil
		return
	}

	item.Previous.Next = item.Next
	item.Next.Previous = item.Previous
	item = nil
}

// Find is a method, which returns pointer of the first quality item
func (list *LinkedList) Find(value int) *Node {

	for currentNode := list.first; currentNode != nil; currentNode = currentNode.Next {
		if currentNode.Value == value {
			return currentNode
		}
	}

	return nil
}
