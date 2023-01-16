package ds

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	curr *Node
	size uint
}

func (mlist *MyLinkedList) Add(val int) {
	newNode := &Node{val, nil}

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
