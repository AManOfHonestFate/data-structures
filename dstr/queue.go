package dstr

import "errors"

var defaultCapacity = 16

type Queue[T any] struct {
	Queue []T
	head  int
	tail  int
	count int
	cap   int
}

type QueueI[T any] interface {
	Push(T)
	Pop() (T, error)
	GetHead() (T, error)
	GetTail() (T, error)
	resize()
}

func NewQueue[T any]() Queue[T] {
	q := make([]T, defaultCapacity)
	return Queue[T]{q, 0, -1, 0, defaultCapacity}
}

func (q *Queue[T]) Push(t T) {
	q.count++
	q.tail++

	if q.tail >= q.cap {
		q.resize()
	}

	q.Queue[q.tail] = t
}

func (q *Queue[T]) Pop() (T, error) {
	var val T
	if q.count > 0 {
		val = q.Queue[q.head]
		q.head++
		q.count--
		return val, nil
	}

	return val, errors.New("Queue is empty")
}

func (q *Queue[T]) GetHead() (T, error) {
	var val T
	if q.count > 0 {
		val = q.Queue[q.head]
		return val, nil
	}

	return val, errors.New("Queue is empty")
}

func (q *Queue[T]) GetTail() (T, error) {
	var val T
	if q.count > 0 {
		val = q.Queue[q.tail]
		return val, nil
	}

	return val, errors.New("Queue is empty")
}

func (q *Queue[T]) resize() {
	if q.count > q.cap {
		q.cap *= 2
	}

	buff := make([]T, q.cap)
	copy(buff, q.Queue)
	q.Queue = buff
	q.head = 0
	q.tail = q.count - 1
}
