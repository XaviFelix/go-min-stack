package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type IntMinHeap []int

// Interface
func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Stack struct {
	stack   []int
	minHeap IntMinHeap
	length  int
}

func NewStack() *Stack {
	s := &Stack{
		stack:   make([]int, 0),
		minHeap: make(IntMinHeap, 0),
		length:  0,
	}
	heap.Init(&s.minHeap)
	return s
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) GetSize() int {
	return s.length
}

func (s *Stack) Push(value int) {
	s.stack = append(s.stack, value)
	heap.Push(&s.minHeap, value)
	s.length++
}

// TODO: NEed to remove element from the heap
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	lastIndex := s.length - 1
	element := s.stack[lastIndex]
	s.stack = s.stack[:lastIndex]
	s.length--

	return element, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s.stack[s.length-1], nil
}

func (s *Stack) GetMin() (int, error) {
	if s.minHeap.Len() == 0 {
		return 0, errors.New("No minimum, heap is empty")
	}

	return s.minHeap[0], nil
}

// Testing the minStack
func main() {
	minStack := NewStack()

	minStack.Push(10)
	minStack.Push(5)
	minStack.Push(20)

	min, err := minStack.GetMin()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Current minimum:", min)
	}

	element, err := minStack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped element:", element)
	}

	min, err = minStack.GetMin()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New minimum:", min)
	}
}
