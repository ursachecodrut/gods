package lists

import (
	"fmt"
)

type Node[T any] struct {
	next *Node[T]
	data T
}

type SinglyLinkedList[T any] struct {
	head *Node[T]
}

func (l SinglyLinkedList[T]) Size() int {
	size := 0
	for l.head != nil {
		size++
		l.head = l.head.next
	}

	return size
}

func (l *SinglyLinkedList[T]) Prepend(data T) {
	newNode := &Node[T]{data: data}
	newNode.next = l.head
	l.head = newNode
}

func (l *SinglyLinkedList[T]) Append(data T) {
	newNode := &Node[T]{data: data}

	if l.head == nil {
		l.head = newNode
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (l *SinglyLinkedList[T]) DeleteAtIndex(index int) {
	if l.head == nil {
		return
	}

	if index == 0 {
		l.head = l.head.next
	}

	temp := l.head
	for i := 0; temp != nil && i < index-1; i++ {
		temp = temp.next
	}

	if temp == nil || temp.next == nil {
		return
	}

	temp.next = temp.next.next
}

func (l SinglyLinkedList[T]) Print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " ")
		temp = temp.next
	}
	fmt.Println()
}
