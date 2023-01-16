package ds

type DNode struct {
	val  int
	next *DNode
	pre  *DNode
}

type MyDoublyLinkedList struct {
	head *DNode
	curr *DNode
	size uint
}

func (mlist *MyDoublyLinkedList) Add(val int) {
	newNode := &Node{val, nil, nil}

	if mlist.size == 0 {
		mlist.head = newNode
		mlist.curr = newNode
	} else {
		mlist.curr.next = newNode
		mlist.curr = newNode
	}
	mlist.size++
}

func (mlist *MyLinkedList) GetHead() *Node {
	headNode := mlist.head
	mlist.head = mlist.head.next
	mlist.size--

	return headNode
}

func (mlist *MyLinkedList) Get() *Node {
	node := mlist.curr

	if mlist.size <= 2 {
		mlist.curr = mlist.head
	} else {
		cnt := uint(0)
		mlist.curr = mlist.head
		for ; cnt < mlist.size-uint(2); cnt++ {
			mlist.curr = mlist.curr.next
			mlist.curr.next = nil
		}
	}
	mlist.curr.next = nil
	mlist.size--

	return node
}
