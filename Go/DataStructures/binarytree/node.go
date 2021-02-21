package binarytree

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
