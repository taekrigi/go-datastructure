package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	length := len(s.items) - 1
	toRemove := s.items[length]
	s.items = s.items[:length]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Pop()
	fmt.Println(myStack)
}