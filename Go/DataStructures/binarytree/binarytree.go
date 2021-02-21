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

// Insert allows to add new item to the tree
func (tree *BinaryTree) Insert(item int) (bool, error) {

	if tree.Root == nil {
		tree.Root = NewNode(item)
		return true, nil
	}

	return tree.Root.Insert(item)
}
