package utils

import "fmt"

func NewStack() Stack {
	return &stack{}
}

type Stack interface {
	Push(data any)
	Pop() (any, error)
	Peek() (any, error)
	IsEmpty() bool
}

type stack struct {
	items []any
}

func (s *stack) Push(data any) {
	s.items = append(s.items, data)
}

func (s *stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(s.items) == 0
}
