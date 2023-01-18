package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	mu sync.Mutex
	s  []int
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]int, 0)}
}

func (st *Stack) isEmpty() bool {
	if len(st.s) == 0 {
		return true
	}
	return false
}

func (st *Stack) push(x int) {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.s = append(st.s, x)
}

func (st *Stack) pop() (int, bool) {
	st.mu.Lock()
	defer st.mu.Unlock()

	l := len(st.s)
	if l == 0 {
		return -1, false
	}

	top := st.s[l-1]
	st.s = st.s[:l-1]

	return top, true
}

func main() {
	s := NewStack()
	fmt.Println("Stack Empty ?", s.isEmpty())
	s.push(11)
	s.push(12)
	s.push(13)
	s.push(17)
	s.push(15)
	fmt.Println("Stack Empty ?", s.isEmpty())
	x, _ := s.pop()
	fmt.Println("Stack Po >", x)
	x, _ = s.pop()
	fmt.Println("Stack Po >", x)
	x, _ = s.pop()
	fmt.Println("Stack Po >", x)

}
