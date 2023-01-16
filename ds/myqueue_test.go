package ds

import (
	"testing"
)

func TestBuildQueue(t *testing.T) {
	capacity := 3
	myQueue := BuildQueue(capacity)

	if capacity != myQueue.Capacity() {
		t.Errorf("Wrong capacity of the queue which except is %d, but got %d\n", capacity, myQueue.Capacity())
	}

	if 0 != myQueue.Size() {
		t.Errorf("Wrong size of the queue which except is %d, but got %d\n", 0, myQueue.Size())
	}

	if !myQueue.IsEmpty() {
		t.Errorf("Test failed since \"IsEmpty\" return false\n")
	}

	if myQueue.IsFull() {
		t.Errorf("Test failed since \"IsFull\" return true\n")
	}
}

func TestEnQueue(t *testing.T) {
	capacity := 3
	entries := make([]QEntry, 3)
	myQueue := MyQueue{capacity, 0, 0, 0, entries}

	if capacity != myQueue.Capacity() {
		t.Errorf("Wrong capacity of the queue which except is %d, but got %d\n", capacity, myQueue.Capacity())
	}

	if 0 != myQueue.Size() {
		t.Errorf("Wrong size of the queue which except is %d, but got %d\n", 0, myQueue.Size())
	}

	if !myQueue.IsEmpty() {
		t.Errorf("Test failed since \"IsEmpty\" return false\n")
	}

	if myQueue.IsFull() {
		t.Errorf("Test failed since \"IsFull\" return true\n")
	}

	myQueue.EnQueue(&QEntry{1})
	myQueue.EnQueue(&QEntry{2})
	myQueue.EnQueue(&QEntry{3})

	if capacity != myQueue.Capacity() {
		t.Errorf("Wrong capacity of the queue which except is %d, but got %d\n", capacity, myQueue.Capacity())
	}

	if capacity != myQueue.Size() {
		t.Errorf("Wrong size of the queue which except is %d, but got %d\n", capacity, myQueue.Size())
	}

	if myQueue.IsEmpty() {
		t.Errorf("Test failed since \"IsEmpty\" return true\n")
	}

	if !myQueue.IsFull() {
		t.Errorf("Test failed since \"IsFull\" return false\n")
	}
}

func TestDeQueue(t *testing.T) {
	capacity := 3
	entries := make([]QEntry, 3)
	entries[0] = QEntry{1}
	entries[1] = QEntry{2}
	entries[2] = QEntry{3}
	myQueue := MyQueue{capacity, capacity, 0, 0, entries}

	if capacity != myQueue.Capacity() {
		t.Errorf("Wrong capacity of the queue which except is %d, but got %d\n", capacity, myQueue.Capacity())
	}

	if capacity != myQueue.Size() {
		t.Errorf("Wrong size of the queue which except is %d, but got %d\n", capacity, myQueue.Size())
	}

	if myQueue.IsEmpty() {
		t.Errorf("Test failed since \"IsEmpty\" return true\n")
	}

	if !myQueue.IsFull() {
		t.Errorf("Test failed since \"IsFull\" return false\n")
	}

	rtnEntry := myQueue.DeQueue()
	if 1 != rtnEntry.GetValue() {
		t.Errorf("Wrong value of the entry which except is %d, but got %d\n", 1, rtnEntry.GetValue())
	}

	rtnEntry = myQueue.DeQueue()
	if 2 != rtnEntry.GetValue() {
		t.Errorf("Wrong value of the entry which except is %d, but got %d\n", 2, rtnEntry.GetValue())
	}

	rtnEntry = myQueue.DeQueue()
	if 3 != rtnEntry.GetValue() {
		t.Errorf("Wrong value of the entry which except is %d, but got %d\n", 3, rtnEntry.GetValue())
	}

	if capacity != myQueue.Capacity() {
		t.Errorf("Wrong capacity of the queue which except is %d, but got %d\n", capacity, myQueue.Capacity())
	}

	if 0 != myQueue.Size() {
		t.Errorf("Wrong size of the queue which except is %d, but got %d\n", 0, myQueue.Size())
	}

	if !myQueue.IsEmpty() {
		t.Errorf("Test failed since \"IsEmpty\" return false\n")
	}

	if myQueue.IsFull() {
		t.Errorf("Test failed since \"IsFull\" return true\n")
	}
}

func TestEnQueueFail(t *testing.T) {
	capacity := 3
	entries := make([]QEntry, 3)
	entries[0] = QEntry{1}
	entries[1] = QEntry{2}
	entries[2] = QEntry{3}
	myQueue := MyQueue{capacity, capacity, 0, 0, entries}

	if capacity != myQueue.Capacity() {
		t.Errorf("Wrong capacity of the queue which except is %d, but got %d\n", capacity, myQueue.Capacity())
	}

	if capacity != myQueue.Size() {
		t.Errorf("Wrong size of the queue which except is %d, but got %d\n", capacity, myQueue.Size())
	}

	if myQueue.IsEmpty() {
		t.Errorf("Test failed since \"IsEmpty\" return true\n")
	}

	if !myQueue.IsFull() {
		t.Errorf("Test failed since \"IsFull\" return false\n")
	}

	isOK := myQueue.EnQueue(&QEntry{4})

	if isOK {
		t.Errorf("Test failed since \"EnQueue\" return true\n")
	}
}

func TestDeQueueFail(t *testing.T) {
	capacity := 3
	entries := make([]QEntry, 3)
	myQueue := MyQueue{capacity, 0, 0, 0, entries}

	if capacity != myQueue.Capacity() {
		t.Errorf("Wrong capacity of the queue which except is %d, but got %d\n", capacity, myQueue.Capacity())
	}

	if 0 != myQueue.Size() {
		t.Errorf("Wrong size of the queue which except is %d, but got %d\n", 0, myQueue.Size())
	}

	if !myQueue.IsEmpty() {
		t.Errorf("Test failed since \"IsEmpty\" return false\n")
	}

	if myQueue.IsFull() {
		t.Errorf("Test failed since \"IsFull\" return true\n")
	}

	//nullEntry := myQueue.DeQueue()
}
