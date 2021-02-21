package binarytree

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
