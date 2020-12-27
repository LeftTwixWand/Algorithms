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

	myHeap.Push(enterprice.NewEnterprice("name1", "address1", time.January, 2001, 1))
	myHeap.Push(enterprice.NewEnterprice("name2", "address2", time.February, 2002, 2))
	myHeap.Push(enterprice.NewEnterprice("name3", "address3", time.March, 2003, 3))
	myHeap.Push(enterprice.NewEnterprice("name4", "address4", time.April, 2004, 4))
	myHeap.Push(enterprice.NewEnterprice("name11", "address11", time.May, 2011, 1))

	myHeap.Print()
}

func main() {

	// myHeap := binaryheap.BinaryHeap{}

	// companies := []enterprice.Enterprice{
	// 	enterprice.Enterprice{Name: "Name1", Address: "Address1", Month: time.Month(1), Year: 1, Profit: 1},
	// 	enterprice.Enterprice{Name: "Name2", Address: "Address2", Month: time.Month(2), Year: 1, Profit: 2},
	// 	enterprice.Enterprice{Name: "Name3", Address: "Address3", Month: time.Month(3), Year: 1, Profit: 3},
	// 	enterprice.Enterprice{Name: "Name4", Address: "Address4", Month: time.Month(4), Year: 1, Profit: 4},
	// }

	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')

	// fmt.Println(text)

	// for _, company := range companies {
	// 	myHeap.Push(company.Profit)
	// }

	// fmt.Println(companies)

	useBinaryHeap()

}
