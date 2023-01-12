package ds

import "testing"

func TestBuildQueue(t *testing.T) {
	capacity := 3
	myQueue := BuildQueue(capacity)

	if capacity != myQueue.Capacity() {
		t.Errorf("Wrong capacity of the queue which except is %d, but got %d\n", capacity, myQueue.Capacity())
	}

	if 0 != myQueue.Size() {
		t.Errorf("Wrong size of the queue which except is %d, but got %d\n", 0, myQueue.Size())
	}
}
