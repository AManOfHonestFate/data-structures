package dstr

import "errors"

var defaultCapacity = 16

type queue[T any] struct {
	queue []T
	head  int
	tail  int
	count int
	cap   int
}

func NewQueue[T any](arr []T) queue[T] {
	q := make([]T, defaultCapacity)
	copy(q, arr)
	return queue[T]{q, len(arr) - 1, 0, len(arr), defaultCapacity}
}

func (q *queue[T]) Push(t T) {
	q.count++
	q.head++

	if q.head >= q.cap {
		q.resize()
	}

	q.queue[q.head] = t
}

func (q *queue[T]) Pop() (T, error) {
	var val T
	if q.count > 0 {
		val = q.queue[q.tail]
		q.tail++
		q.count--
		return val, nil
	}

	return val, errors.New("queue is empty")
}

func (q *queue[T]) GetHead() (T, error) {
	var val T
	if q.count > 0 {
		val = q.queue[q.head]
		return val, nil
	}

	return val, errors.New("queue is empty")
}

func (q *queue[T]) GetTail() (T, error) {
	var val T
	if q.count > 0 {
		val = q.queue[q.tail]
		return val, nil
	}

	return val, errors.New("queue is empty")
}

func (q *queue[T]) resize() {
	if q.count > q.cap {
		q.cap *= 2
	}

	buff := make([]T, q.cap)
	copy(buff, q.queue)
	q.queue = buff
	q.tail = 0
	q.head = q.count - 1
}

type Queue[T any] interface {
	Push(T)
	Pop() (T, error)
	GetHead() (T, error)
	GetTail() (T, error)
	resize()
}
