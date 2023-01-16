package ds

import (
	"testing"
)

func TestBuildStack(t *testing.T) {
	capacity := 10
	myStack := BuildStack(capacity)

	if capacity != myStack.GetCapacity() {
		t.Errorf("Wrong capacity that except is %d but got %d\n", capacity, myStack.GetCapacity())
	}

	if 0 != myStack.GetSize() {
		t.Errorf("Wrong size that except is %d but got %d\n", 0, myStack.GetSize())
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

	if 3 != myStack.GetSize() {
		t.Errorf("Wrong size of stack that except is %d but got %d\n", 3, myStack.GetSize())
	}
}

func TestPopFromStack(t *testing.T) {
	_capacity := 3
	_entries := []Entry{Entry{Val: 1}, Entry{Val: 2}, Entry{Val: 3}}
	myStack := MyStack{capacity: _capacity, currSize: _capacity, entries: _entries}

	if _capacity != myStack.GetCapacity() {
		t.Errorf("Wrong capacity that except is %d but got %d\n", _capacity, myStack.GetCapacity())
	}

	if _capacity != myStack.GetSize() {
		t.Errorf("Wrong size that except is %d but got %d\n", 3, myStack.GetSize())
	}

	rtnVal := myStack.Pop().Val
	if 3 != rtnVal {
		t.Errorf("Wrong value from pop() which except is %d but got %d\n", 3, rtnVal)
	}

	rtnVal = myStack.Pop().Val
	if 2 != rtnVal {
		t.Errorf("Wrong value from pop() which except is %d but got %d\n", 2, rtnVal)
	}

	rtnVal = myStack.Pop().Val
	if 1 != rtnVal {
		t.Errorf("Wrong value from pop() which except is %d but got %d\n", 1, rtnVal)
	}

	if _capacity != myStack.GetCapacity() {
		t.Errorf("Wrong capacity that except is %d but got %d\n", _capacity, myStack.GetCapacity())
	}

	if 0 != myStack.GetSize() {
		t.Errorf("Wrong size that except is %d but got %d\n", 0, myStack.GetSize())
	}
}

func TestPushExceedCapacity(t *testing.T) {
	entry1 := Entry{Val: 1}
	entry2 := Entry{Val: 2}
	entry3 := Entry{Val: 3}
	entry4 := Entry{Val: 4}

	myStack := BuildStack(3)

	if !myStack.Push(entry1) {
		t.Errorf("Push Entry with value %d failed\n", entry1.Val)
	}
	if !myStack.Push(entry2) {
		t.Errorf("Push Entry with value %d failed\n", entry2.Val)
	}
	if !myStack.Push(entry3) {
		t.Errorf("Push Entry with value %d failed\n", entry3.Val)
	}
	if myStack.Push(entry4) {
		t.Errorf("Push Entry with value %d success but it should be fail since exceed the capacity of the stack\n", entry4.Val)
	}

}

func TestPopFromEmptyStack(t *testing.T) {
	_capacity := 3
	_entries := []Entry{Entry{Val: 1}, Entry{Val: 2}, Entry{Val: 3}}
	myStack := MyStack{capacity: _capacity, currSize: _capacity, entries: _entries}

	if _capacity != myStack.GetCapacity() {
		t.Errorf("Wrong capacity that except is %d but got %d\n", _capacity, myStack.GetCapacity())
	}

	if _capacity != myStack.GetSize() {
		t.Errorf("Wrong size that except is %d but got %d\n", 3, myStack.GetSize())
	}

	rtnVal := myStack.Pop().Val
	if 3 != rtnVal {
		t.Errorf("Wrong value from pop() which except is %d but got %d\n", 3, rtnVal)
	}

	rtnVal = myStack.Pop().Val
	if 2 != rtnVal {
		t.Errorf("Wrong value from pop() which except is %d but got %d\n", 2, rtnVal)
	}

	rtnVal = myStack.Pop().Val
	if 1 != rtnVal {
		t.Errorf("Wrong value from pop() which except is %d but got %d\n", 1, rtnVal)
	}

	if _capacity != myStack.GetCapacity() {
		t.Errorf("Wrong capacity that except is %d but got %d\n", _capacity, myStack.GetCapacity())
	}

	if 0 != myStack.GetSize() {
		t.Errorf("Wrong size that except is %d but got %d\n", 0, myStack.GetSize())
	}

	rtnVal = myStack.Pop().Val
	if 0 != rtnVal {
		t.Errorf("Wrong value from pop() which except is %d but got %d\n", 0, rtnVal)
	}
}
