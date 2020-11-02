package node

// Node is a struct, which implements item in linked list
type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

// NewNode is a constructor, which returns pointer on new Node struct
func NewNode(value int) *Node {
	return &Node{Value: value}
}
