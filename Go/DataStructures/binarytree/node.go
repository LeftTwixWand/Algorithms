package binarytree

import (
	"errors"
	"strconv"
	"strings"
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

	if node == nil {
		(*markup)[level] += "+.. ..+  "

		return
	}

	if node.Left != nil {
		(*markup)[level] += "+--"
	}
	(*markup)[level] += strconv.Itoa(node.Value)
	if node.Right != nil {
		(*markup)[level] += "--+"
	}
	(*markup)[level] += "     "

	node.Left.PrintNode(markup, level+1)
	node.Right.PrintNode(markup, level+1)
}

// |..data..| - mask, value - number, capacity - value capacity
func applyMask(mask string, value, capacity int) string {

	return strings.Replace(mask, "data", strconv.Itoa(value), 0)
}

func insertAt(basic, value string, pos int) string {
	return basic[:pos] + value + basic[pos:]
}

// func (node *Node) PrintNode(markup *[]string, level int) {

// 	if len(*markup) < level+1 {
// 		*markup = append(*markup, "")
// 	}

// 	if node.Left != nil {
// 		(*markup)[level] += "+--"

// 		node.Left.PrintNode(markup, level+1)
// 	} else {
// 		(*markup)[level] += "   "
// 	}

// 	(*markup)[level] += strconv.Itoa(node.Value)

// 	if node.Right != nil {
// 		(*markup)[level] += "--+"

// 		node.Right.PrintNode(markup, level+1)
// 	} else {
// 		(*markup)[level] += "   "
// 	}

// 	// *markup += "\n" + strconv.Itoa(node.Value)
// }
