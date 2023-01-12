package ds

import "fmt"

type QEntry struct {
	val int
}

func (e *QEntry) GetValue() int {
	return e.val
}

type MyQueue struct {
	capacity int
	currSize int
	headIdx  int
	tailIdx  int
	entries  []QEntry
}

func BuildQueue(_capacity int) *MyQueue {
	_entries := make([]QEntry, _capacity)
	myQueue := MyQueue{capacity: _capacity, currSize: 0, headIdx: 0, tailIdx: 0, entries: _entries}
	return &myQueue
}

func (q *MyQueue) IsEmpty() bool {
	return q.currSize == 0
}

func (q *MyQueue) IsFull() bool {
	return q.capacity == q.currSize
}

func (q *MyQueue) Capacity() int {
	return q.capacity
}

func (q *MyQueue) Size() int {
	return q.currSize
}

func (q *MyQueue) EnQueue(e *QEntry) bool {
	fmt.Printf("Invoke enQueue\n")
	if !q.IsFull() {
		q.entries[q.headIdx] = *e

		q.checkAndReindexOfHead()
		q.headIdx++
		q.currSize++
		return true
	}
	return false
}

func (q *MyQueue) DeQueue() *QEntry {
	fmt.Printf("Invoke deQueue\n")

	if !q.IsEmpty() {
		rtnEntry := q.entries[q.tailIdx]

		q.checkAndReindexOfTail()
		q.tailIdx++
		q.currSize--
		return &rtnEntry
	}
	return &QEntry{}
}

func (q *MyQueue) checkAndReindexOfHead() {
	if q.headIdx == (q.capacity - 1) {
		q.headIdx = 0
	}
}

func (q *MyQueue) checkAndReindexOfTail() {
	if q.tailIdx == (q.capacity - 1) {
		q.tailIdx = 0
	}
}
