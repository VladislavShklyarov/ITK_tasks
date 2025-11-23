package main

type Stacker interface {
	Push(v int)
	Pop() int
}

type Stack struct {
	data []int
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)

}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	lastIndex := len(s.data) - 1
	v := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	return v
}

func New() *Stack {
	return &Stack{data: make([]int, 0)}
}
