package binarytree

import (
	"errors"
)

// Node is a data structure for the binatytree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewNode is the Node constructor
func NewNode(data int) *Node {
	return &Node{
		Value: data,
	}
}

// Insert is a mathod, which allows to insert a new value to the tree
func (node *Node) Insert(item int) (bool, error) {

	if node.Value == item {
		return false, errors.New("This value is already exist in a tree")
	}

	if node.Value > item {

		if node.Left == nil {
			node.Left = NewNode(item)
			return true, nil
		}

		return node.Left.Insert(item)
	}

	if node.Value < item {

		if node.Right == nil {
			node.Right = NewNode(item)
			return true, nil
		}

		return node.Right.Insert(item)
	}

	return false, errors.New("Error in processing")
}
