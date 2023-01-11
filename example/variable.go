package main

import "fmt"

var a, b, c bool

func showVariable() {
	var i, j, k int
	// %t the word true or false
	// %v the value in a default format
	fmt.Printf("%d, %d, %t, %t, %v, %d\n", i, j, a, b, c, k)
}
