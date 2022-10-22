package dstr

import "testing"

func TestQueue_GetTail(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.GetTail()
	if err == nil {
		t.Error("Empty queue don't error on GetTail")
	}

	tail := 19
	for i := 0; i <= tail; i++ {
		q.Push(i)
	}
	val, err := q.GetTail()
	if val != tail {
		t.Error("Wrong tail value: ", q)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestQueue_GetHead(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.GetHead()
	if err == nil {
		t.Error("Empty queue don't error on GetHead")
	}

	for i := 1; i < 20; i++ {
		q.Push(i)
	}
	val, err := q.GetHead()
	if val != 1 {
		t.Error("Wrong head value: ", q)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestPushAndPop(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 70; i++ {
		q.Push(i)
	}

	for i := 0; i < 70; i++ {
		val, err := q.Pop()

		if err != nil {
			t.Error(err)
		}

		if val != i {
			t.Fatalf("Wrong value, Pushed %v, Poped %v", i, val)
		}
	}

	_, err := q.Pop()
	if err == nil {
		t.Error("Successfully Popped empty queue")
	}
}
