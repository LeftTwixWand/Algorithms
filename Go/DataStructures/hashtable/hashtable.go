package hashtable

import (
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

	if item != nil {
		return item.Value
	}

	return nil
}

func (table *HashTable) getNodeByHash(hash int) *KeyValuePair {
	return table.data[hash]
}

// hash is a method, which returns a hash for the specific key
func (table *HashTable) hash(key int) int {
	return (key % table.Capacity) / 100
}

// Print is a test method for printion all the table
func (table *HashTable) Print() {

	// for _, item := range table.data {
	// 	fmt.Println("Key = ", item.Key, "Value = ", item.Value)
	// }

	const padding = 3
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(writer, "a\tb\taligned\t")
	fmt.Fprintln(writer, "aa\tbb\taligned\t")
	fmt.Fprintln(writer, "aaa\tbbb\tunaligned") // no trailing tab
	fmt.Fprintln(writer, "aaaa\tbbbb\taligned\t")
	writer.Flush()
}
