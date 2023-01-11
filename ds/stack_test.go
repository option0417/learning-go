package ds

import "testing"

func TestBuildStack(t *testing.T) {
	capacity := 10
	myStack := BuildStack(capacity)

	if capacity != myStack.GetCapacity() {
		t.Errorf("Wrong capacity that except is %d but got %d\n", capacity, myStack.GetCapacity())
	}

	if 0 != myStack.GetSize() {
		t.Errorf("Wrong size that except is 0 but got %d\n", myStack.GetSize())
	}
}

func TestPushToStack(t *testing.T) {
	entry1 := Entry{Val: 1}
	entry2 := Entry{Val: 2}
	entry3 := Entry{Val: 3}

	myStack := BuildStack(5)

	if !myStack.Push(entry1) {
		t.Errorf("Push Entry with value %d failed\n", entry1.Val)
	}
	if !myStack.Push(entry2) {
		t.Errorf("Push Entry with value %d failed\n", entry2.Val)
	}
	if !myStack.Push(entry3) {
		t.Errorf("Push Entry with value %d failed\n", entry3.Val)
	}
}
