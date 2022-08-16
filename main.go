package main

import (
	"fmt"

	Lists "github.com/ursachecodrut/gods/lists"
)

func main() {
	var l Lists.SinglyLinkedList[int]
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	l.DeleteAtIndex(10)
	l.Print()
	fmt.Println(l.Size())
}
