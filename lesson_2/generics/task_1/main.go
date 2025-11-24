package main

import (
	"fmt"
)

// решил вместо классического стэка реализовать на базе связанного списка

type Stack[T any] struct {
	head *StackNode[T]
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

type StackNode[T any] struct {
	value T
	next  *StackNode[T]
}

func (s *Stack[T]) Push(value T) {

	newNode := &StackNode[T]{
		value: value,
		next:  nil,
	}

	if s.IsEmpty() {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}

	s.size++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	value := s.head.value
	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack[T]) Print() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}
	curr := s.head
	fmt.Print("[")
	for curr != nil {
		fmt.Printf("(%v)", curr.value)
		if curr.next != nil {
			fmt.Printf(" -> ")
		}
		curr = curr.next
	}
	fmt.Print("]")
	fmt.Println()
}

func main() {
	stack := NewStack[string]()
	stack.Push("Alfa")
	stack.Push("Beta")
	stack.Push("Gamma")

	stack.Print()
	fmt.Println(stack.Pop())
	stack.Print()

	fmt.Println(stack.Peek())

	fmt.Println(stack.IsEmpty())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Print()
}
