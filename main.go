package main

import (
	"data-structures/dstr"
	"fmt"
)

func main() {
	q := dstr.NewQueue([]int{0, 1})
	for i := 2; i < 70; i++ {
		q.Push(i)
	}
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
