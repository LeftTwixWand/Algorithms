package main

import (
	"fmt"

	"./hashtable"
)

func main() {
	var table = hashtable.New()

	table.Print()
	fmt.Println("Hello, World!")
}
