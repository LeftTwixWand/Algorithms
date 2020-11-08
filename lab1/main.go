package main

import (
	"fmt"
	"time"

	"./binaryheap"
	"./enterprice"
	"./linkedlist"
)

func useLinkedList() {

	myList := linkedlist.LinkedList{}

	myList.PushBackDouble(2)
	item := myList.PushBackDouble(3)
	myList.PushFront(1)
	myList.InsertAt(item, 4)
	myList.Print()
	fmt.Println("Count items = ", myList.Count)

	myList.DeleteAt(item.Next)
	myList.Print()
	fmt.Println("Count items = ", myList.Count)
}

func useBinaryHeap() {

	myHeap := binaryheap.BinaryHeap{}

	myHeap.Push(2)
	myHeap.Push(1)
	myHeap.Push(3)
	myHeap.Push(3)
	myHeap.Push(2)

	myHeap.Print()
}

func main() {

	myHeap := binaryheap.BinaryHeap{}

	companies := []enterprice.Enterprice{
		enterprice.Enterprice{Name: "Name1", Address: "Address1", Month: time.Month(1), Year: 1, Profit: 1},
		enterprice.Enterprice{Name: "Name2", Address: "Address2", Month: time.Month(2), Year: 1, Profit: 2},
		enterprice.Enterprice{Name: "Name3", Address: "Address3", Month: time.Month(3), Year: 1, Profit: 3},
		enterprice.Enterprice{Name: "Name4", Address: "Address4", Month: time.Month(4), Year: 1, Profit: 4},
	}

	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')

	// fmt.Println(text)

	for _, company := range companies {
		myHeap.Push(company.Profit)
	}

	fmt.Println(companies)

}
