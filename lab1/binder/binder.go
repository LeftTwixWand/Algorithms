package binder

// Binder is a node in binery heap structure
type Binder struct {
	Value               int
	parent, left, right *Binder
}

// NewBinder returns a pointer on a new Binder structure
func NewBinder(value int) *Binder {
	return &Binder{Value: value}
}
