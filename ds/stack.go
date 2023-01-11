package ds

import "fmt"

type MyStack struct {
	capacity int
	currSize int
	entries  []Entry
}

type Entry struct {
	Val int
}

func BuildMyStack(capacity int) *MyStack {
	entries := make([]Entry, capacity)
	myStack := MyStack{capacity, 0, entries}
	return &myStack
}

func (s *MyStack) Push(entry Entry) bool {
	fmt.Printf("Invoke push\n")

	if !s.isFull() {
		s.entries[s.currSize] = entry
		s.currSize++
		return true
	} else {
		return false
	}
}

func (s *MyStack) Pop() Entry {
	fmt.Printf("Invoke pop\n")

	var rtnEntry Entry

	if !s.isEmpty() {
		s.currSize--
		rtnEntry = s.entries[s.currSize]
	}

	return rtnEntry
}

func (s *MyStack) isFull() bool {
	return s.capacity == s.currSize
}

func (s *MyStack) isEmpty() bool {
	return s.currSize == 0
}
