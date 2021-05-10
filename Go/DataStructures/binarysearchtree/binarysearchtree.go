package binarysearchtree

import "fmt"

func Use() {

	tree := &Node{CarNumber: 100}

	tree.insert(50)

	fmt.Println(tree)

}

func (node *Node) insert(key int) {
	if key < node.CarNumber {
		if node.Right == nil {
			node.Right = &Node{CarNumber: key}
		} else {
			node.Right.insert(key)
		}
	} else if node.CarNumber > key {
		if node.Left == nil {
			node.Left = &Node{CarNumber: key}
		} else {
			node.Left.insert(key)
		}
	}
}

func (node *Node) search(key int) bool {

	if node == nil {
		return false
	}

	if node.CarNumber < key {
		return node.Right.search(key)
	} else if node.CarNumber > key {
		return node.Left.search(key)
	}

	return true
}
