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

	prefix := buildPrefix(level)
	// postfix := calculatePostfix(level)
	(*markup)[level] += prefix + strconv.Itoa(node.Value) + reverse(prefix) + " "

	node.Left.PrintNode(markup, level+1)
	node.Right.PrintNode(markup, level+1)
}

func buildPrefix(level int) string {

	// A format has the next structure:
	// lvl2   |   +---k---+
	// lvl1   | +-k-+   +-k-+
	// lvl0   |+k+ +k+ +k+ +k+
	// lvl -1 |k k k k k k k k

	// So, prefix has the next structure: lvl * 2 + 1

	result := ""

	indent := calculateFullPrefixIndent(level)
	centralIndex := getCentralIndex(indent)

	result += getPrefixWhiteSpaces(centralIndex)

	if indent > 0 {
		result += "+"
	}

	result += getPrefixRelations(centralIndex)

	return result
}

func reverse(text string) string {

	result := ""

	textLength := len(text)

	for i := textLength; i > 0; i-- {
		result += string(text[i])
	}

	return result
}

func getCentralIndex(number int) int {
	return (number-1)/2 + 1
}

func calculateFullPrefixIndent(level int) int {

	if level == 0 {
		return 0
	}

	return calculateFullPrefixIndent(level-1)*2 + 1
}

func getPrefixWhiteSpaces(centralIndex int) string {

	prefix := ""

	for i := 0; i < centralIndex-1; i++ {
		prefix += " "
	}

	return prefix
}

func getPrefixRelations(centralIndex int) string {
	prefix := ""

	for i := 0; i < centralIndex-1; i++ {
		prefix += "-"
	}

	return prefix
}
