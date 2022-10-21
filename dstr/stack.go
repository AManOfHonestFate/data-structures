package dstr

import "errors"

type Stack[T any] struct {
	stack []T
}

func NewStack[T any](arr []T) Stack[T] {
	return Stack[T]{arr}
}

func (s *Stack[T]) Push(newElem T) {
	s.stack = append(s.stack, newElem)
}

func (s *Stack[T]) Pop() (T, error) {
	elem, err := s.Head()

	if err != nil {
		return elem, err
	}

	s.stack = s.stack[:len(s.stack)-1]

	return elem, nil
}

func (s *Stack[T]) Head() (T, error) {
	var res T
	l := len(s.stack)

	if l == 0 {
		return res, errors.New("stack is empty")
	}

	res = s.stack[l-1]
	return res, nil
}

func (s *Stack[T]) AsArray() []T {
	return s.stack
}

type IStack[T any] interface {
	Push(newElem T)
	Pop() (T, error)
	Head() (T, error)
	AsArray() []T
}
