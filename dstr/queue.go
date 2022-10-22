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

func NewQueue[T any]() queue[T] {
	q := make([]T, defaultCapacity)
	return queue[T]{q, 0, -1, 0, defaultCapacity}
}

func (q *queue[T]) Push(t T) {
	q.count++
	q.tail++

	if q.tail >= q.cap {
		q.resize()
	}

	q.queue[q.tail] = t
}

func (q *queue[T]) Pop() (T, error) {
	var val T
	if q.count > 0 {
		val = q.queue[q.head]
		q.head++
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
	q.head = 0
	q.tail = q.count - 1
}

type Queue[T any] interface {
	Push(T)
	Pop() (T, error)
	GetHead() (T, error)
	GetTail() (T, error)
	resize()
}
