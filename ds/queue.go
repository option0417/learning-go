package ds

import "fmt"

type Entry struct {
	val int
}

func (e *Entry) GetValue() int {
	return e.val
}

type MyQueue struct {
	capacity int
	currSize int
	headIdx  *Entry
	tailIdx  *Entry
	entries  []Entry
}

func (q *MyQueue) isEmpty() bool {
	return q.currSize == 0
}

func (q *MyQueue) isFull() bool {
	return q.capacity == q.currSize
}

func (q *MyQueue) enQueue(e *Entry) bool {
	fmt.Printf("Invoke enQueue\n")
	if !isFull() {
		q.currSize++
	}
}

func (q *MyQueue) deQueue() *Entry {
	fmt.Printf("Invoke deQueue\n")

	if !q.isEmpty() {
		q.currSize--
	}
}
