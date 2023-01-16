package ds

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	myLinkedList := MyLinkedList{}

	myLinkedList.Add(1)
	myLinkedList.Add(2)
	myLinkedList.Add(3)

	if 3 != myLinkedList.size {
		t.Errorf("Wrong size of LinkedList that except is %d but got %d\n", 3, myLinkedList.size)
	}

	headNode := myLinkedList.head
	if 1 != headNode.val {
		t.Errorf("Wrong value of the head-node that except is %d but got %d\n", 1, headNode.val)
	}

	if 2 != headNode.next.val {
		t.Errorf("Wrong value of the next node of the head node that except is %d but got %d\n", 2, headNode.next.val)
	}

	currNode := myLinkedList.curr
	if 3 != currNode.val {
		t.Errorf("Wrong value of the curr-node that except is %d but got %d\n", 3, currNode.val)
	}
}

func TestGetNode(t *testing.T) {
	node1 := &Node{1, nil}
	node2 := &Node{2, nil}
	node3 := &Node{3, nil}

	node1.next = node2
	node2.next = node3

	myLinkedList := MyLinkedList{}
	myLinkedList.head = node1
	myLinkedList.curr = node3
	myLinkedList.size = uint(3)

	rtnNode := myLinkedList.Get()
	if 3 != rtnNode.val {
		t.Errorf("Wrong value of the node which except is %d but got %d\n", 3, rtnNode.val)
	}

	if nil != rtnNode.next {
		t.Errorf("Test failed since the pointer of the next node is not \"nil\", is %v\n", rtnNode.next)
	}

	rtnNode = myLinkedList.Get()
	if 2 != rtnNode.val {
		t.Errorf("Wrong value of the node which except is %d but got %d\n", 2, rtnNode.val)
	}

	if nil != rtnNode.next {
		t.Errorf("Test failed since the pointer of the next node is not \"nil\", is %v\n", rtnNode.next)
	}

	rtnNode = myLinkedList.Get()
	if 1 != rtnNode.val {
		t.Errorf("Wrong value of the node which except is %d but got %d\n", 1, rtnNode.val)
	}

	if nil != rtnNode.next {
		t.Errorf("Test failed since the pointer of the next node is not \"nil\", is %v\n", rtnNode.next)
	}
}
