package dstr

type node[T any] struct {
	next, previous *node[T]
	value          T
}

type List[T any] struct {
	head, tail *node[T]
}

type ListI[T any] interface {
	NewList() *List[T]
}
