package binarytree

import (
	"fmt"
)

// BinaryTree is a data structure
type BinaryTree struct {
	Root *Node
}

// New is a constructor for BinaryTree
func New(root *Node) *BinaryTree {
	return &BinaryTree{
		Root: root,
	}
}

// Insert allows to add new item to the tree
func (tree *BinaryTree) Insert(item int) (bool, error) {

	if tree.Root == nil {
		tree.Root = NewNode(item)
		return true, nil
	}

	return tree.Root.Insert(item)
}

// Print tree to console
func (tree *BinaryTree) Print() {

	if tree.Root == nil {
		fmt.Println("Tree is empty!")
		return
	}

	var markup []string
	tree.Root.PrintNode(&markup, 0)

	for _, item := range markup {
		fmt.Println(item)
	}
}
