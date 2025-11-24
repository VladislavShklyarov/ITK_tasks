package main

import (
	"testing"
)

func TestStack_PushPeekPop(t *testing.T) {
	tests := []struct {
		name       string
		pushValues []int
		wantPeek   int
		wantSize   int
	}{
		{
			name:       "push one value",
			pushValues: []int{10},
			wantPeek:   10,
			wantSize:   1,
		},
		{
			name:       "push multiple values",
			pushValues: []int{1, 2, 3},
			wantPeek:   3,
			wantSize:   3,
		},
		{
			name:       "push nothing",
			pushValues: []int{},
			wantSize:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int]()

			for _, v := range tt.pushValues {
				stack.Push(v)
			}

			if stack.size != tt.wantSize {
				t.Fatalf("size = %d, want %d", stack.size, tt.wantSize)
			}

			if tt.wantSize == 0 {
				return
			}

			got, ok := stack.Peek()
			if !ok {
				t.Fatal("Peek returned !ok, expected a value")
			}

			if got != tt.wantPeek {
				t.Fatalf("Peek() = %d, want %d", got, tt.wantPeek)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()

	// пустой стек
	if _, ok := stack.Pop(); ok {
		t.Fatalf("Pop() returned ok on empty stack, expected false")
	}

	// кладём 1,2,3 и снимаем в обратном порядке
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	testOrder := []int{3, 2, 1}
	for _, want := range testOrder {
		got, ok := stack.Pop()
		if !ok {
			t.Fatalf("Pop returned !ok but stack has values")
		}
		if got != want {
			t.Fatalf("Pop() = %d, want %d", got, want)
		}
	}

	// теперь стек пустой
	if !stack.IsEmpty() {
		t.Fatalf("stack should be empty after removing all elements")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack[int]()

	if !stack.IsEmpty() {
		t.Fatalf("new stack should be empty")
	}

	stack.Push(10)

	if stack.IsEmpty() {
		t.Fatalf("stack should not be empty after Push")
	}

	stack.Pop()

	if !stack.IsEmpty() {
		t.Fatalf("stack should be empty after Pop")
	}
}
