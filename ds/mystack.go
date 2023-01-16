package ds

type MyStack struct {
	capacity int
	currSize int
	entries  []Entry
}

type Entry struct {
	Val int
}

func BuildStack(capacity int) *MyStack {
	entries := make([]Entry, capacity)
	myStack := MyStack{capacity, 0, entries}
	return &myStack
}

func (s *MyStack) Push(entry Entry) bool {
	if !s.isFull() {
		s.entries[s.currSize] = entry
		s.currSize++
		return true
	} else {
		return false
	}
}

func (s *MyStack) Pop() Entry {
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

func (s *MyStack) GetCapacity() int {
	return s.capacity
}

func (s *MyStack) GetSize() int {
	return s.currSize
}
