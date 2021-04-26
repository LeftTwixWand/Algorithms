package binarytree

import (
	"errors"
	"strconv"
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

// PrintNode allows to print tree by root node
func (node *Node) PrintNode(markup *[]string, level int) {

	if len(*markup) < level+1 {
		*markup = append(*markup, "")
	}

	prefix := calculatePrefix(level)
	postfix := calculatePostfix(level)
	(*markup)[level] += prefix + strconv.Itoa(node.Value) + postfix

	node.Left.PrintNode(markup, level+1)
	node.Right.PrintNode(markup, level+1)
}

func calculatePrefix(level int) string {

	result := ""

	for i := 0; i < level; i++ {

		if i%2 == 1 || i == 1 {
			result += "+"
			continue
		}

		if i > level/2 {
			result += "-"
			continue
		}

		result += " "
	}

	return result
}

func calculatePostfix(level int) string {

	result := ""

	for i := 0; i < level; i++ {
		if i == 0 {

		}
	}

	return result
}
