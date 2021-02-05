package hashtable

import (
	"fmt"

	"../keyvaluepair"
)

// HashTable is a data structure
type HashTable struct {
	table []*keyvaluepair.KeyValuePair
}

// New is a constructor for the HashTable structure
func New() *HashTable {
	return &HashTable{}
}

// Print is a test method
func (table *HashTable) Print() {
	fmt.Println("Works!")
}
