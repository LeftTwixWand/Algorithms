package enterprice

import "time"

// Enterprice is a structure, which allows us to store informatqion about some company
type Enterprice struct {
	Name    string
	Address string
	Month   time.Month
	Year    int
	Profit  int
}

// NewEnterprice is a constructor, which returns a pointer of new enterprice structure
func NewEnterprice(name, address string, month time.Month, year, profit int) *Enterprice {
	return &Enterprice{
		Name:    name,
		Address: address,
		Month:   month,
		Year:    year,
		Profit:  profit,
	}
}
