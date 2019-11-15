package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		fmt.Printf("%v\n", (*t).Value)
		Walk((*t).Left, ch)

		//ch <- (*t).Value

		Walk((*t).Right, ch)
	} else {
		fmt.Printf("Nil\n")
	}
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	ch := make(chan int)

	t := tree.New(1)
	fmt.Printf("Tree: %s\n", t.String())
	go Walk(t, ch)

	time.Sleep(3000 * time.Millisecond)
	//for i := range ch {
	//	fmt.Println(i)
	//}
}
