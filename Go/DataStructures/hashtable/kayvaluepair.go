package hashtable

// KeyValuePair is a data structude for a Chaining Hash Table
type KeyValuePair struct {
	Key   string
	Value interface{}
	Next  *KeyValuePair
}

// NewKeyValuePair is a constructor for the KeyValuePair structure
func NewKeyValuePair(key string, value interface{}) *KeyValuePair {
	return &KeyValuePair{Key: key, Value: value}
}
