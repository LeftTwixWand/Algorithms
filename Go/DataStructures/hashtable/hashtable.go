package hashtable

import (
	"fmt"
)

// HashTable is a data structure
type HashTable struct {
	data []*KeyValuePair
}

// New is a constructor for the HashTable structure
func New(size int) *HashTable {
	return &HashTable{
		data: make([]*KeyValuePair, size),
	}
}

// Get returns a value of a specific key
func (table *HashTable) Get(key interface{}) interface{} {
	item := table.getNodeByHash(table.hash(key))

	if item != nil {
		return item.Value
	}

	return nil
}

func (table *HashTable) getNodeByHash(hash uint64) *KeyValuePair {
	return table.data[int(hash)]
}

// hash is a method, which returns a hash for the specific key
func (table *HashTable) hash(key interface{}) uint64 {
	return 1
}

// Print is a test method for printion all the table
func (table *HashTable) Print() {

	for _, item := range table.data {
		fmt.Println("Key = ", item.Key, "Value = ", item.Value)
	}
}
