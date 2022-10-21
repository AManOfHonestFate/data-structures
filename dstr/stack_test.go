package dstr

import "testing"

func TestStack_Head(t *testing.T) {
	st1 := Stack[int]{stack: []int{1, 2}}
	head1, err := st1.Head()
	if head1 != 2 {
		t.Error("Head has wrong value, expecting 2, got ", head1)
	}

	if err != nil {
		t.Error("Not Empty stack has no head")
	}

	st2 := Stack[string]{}
	_, err = st2.Head()
	if err == nil {
		t.Error("Empty stack has head")
	}
}

func TestStack_Pop(t *testing.T) {
	st1 := Stack[bool]{stack: []bool{true, false, true}}
	val, err := st1.Pop()

	if val != true {
		t.Error("Wrong value, expecting true, got ", val)
	}

	if err != nil {
		t.Error("Not Empty Stack can't pop")
	}

	st2 := Stack[float64]{}
	_, err = st2.Pop()

	if err == nil {
		t.Error("Empty stack can pop")
	}
}

func TestStack_Push(t *testing.T) {
	type Person struct {
		name    string
		address int16
	}
	p1 := Person{"Egor", 1040}
	p2 := Person{"Vadim", 3020}

	st := Stack[Person]{}
	st.Push(p1)

	p, _ := st.Head()
	if p != p1 {
		t.Error("Wrong value was pushed")
	}

	st.Push(p2)
	p, _ = st.Head()
	if p != p2 {
		t.Error("Wrong value was pushed")
	}
}
