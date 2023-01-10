package ds

import "fmt"

type MyStack struct {
	capacity int
	currSize int

	entries []Entry
}

type Entry struct {
	val int
}

func BuildMyStack(capacity int) *MyStack {
	entries := make([]Entry, capacity)
	myStack := MyStack{capacity, 0, entries}
	return &myStack
}

func (s MyStack) push(entry Entry) bool {
	fmt.Printf("Invoke push\n")
	return false
}

func (s MyStack) pop() bool {
	fmt.Printf("Invoke pop\n")
	return false
}
