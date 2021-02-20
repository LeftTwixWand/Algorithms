package hashtable

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
)

// HashTable is a data structure
type HashTable struct {
	data     []*KeyValuePair
	Capacity int
}

// New is a constructor for the HashTable structure
func New(size int) *HashTable {
	return &HashTable{
		data:     make([]*KeyValuePair, size),
		Capacity: size,
	}
}

// Get returns a value of a specific key
func (table *HashTable) Get(key int) interface{} {
	item := table.getNodeByHash(table.hash(key))

	return item.Value
}

// Add is a method to and a new item to the table
func (table *HashTable) Add(item *KeyValuePair) (bool, error) {

	if item == nil {
		return false, errors.New("KeyValuePair is nil")
	}

	hash := table.hash(item.Key)

	if table.data[hash] == nil {
		table.data[hash] = item
		return true, nil
	}

	for i := hash + 1; i < len(table.data); i++ {

		if table.data[i] == nil {
			table.data[i] = item
			return true, nil
		}
	}

	for i := 0; i < hash; i++ {

		if table.data[i] == nil {
			table.data[i] = item
			return true, nil
		}
	}

	return false, errors.New("Array length limit error")

}

func (table *HashTable) getNodeByHash(hash int) *KeyValuePair {
	return table.data[hash]
}

// hash is a method, which returns a hash for the specific key
func (table *HashTable) hash(key int) int {
	return key % table.Capacity
}

// Print is a test method for printion all the table
func (table *HashTable) Print() {

	const padding = 3
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)

	for _, item := range table.data {

		if item == nil {
			continue
		}

		// fmt.Println("Key = ", item.Key, "Value = ", item.Value)

		fmt.Fprintln(writer, "Key: ", item.Key, "\t Value: ", item.Value)
	}

	writer.Flush()
}
