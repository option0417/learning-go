package main

import "fmt"

func main() {
  // Say Hello, World
  hello()

  // Define function
  fmt.Printf("Add 1, 2 = %d\n", add2(1, 2))
  fmt.Printf("Add 3, 4, 5 = %d\n", add3(3, 4, 5))

  // Return multiple result
  a := "hello"
  b := "world"
  fmt.Printf("Origin: %s, %s\t", a, b)
  a, b = string_swap(a, b)
  fmt.Printf("Swap: %s, %s\n", a, b)

  // Return named result
  fmt.Printf("%d remainder 9 = %d\n", 12, remainder9(12))

  // Define list of variables
  showVariable()
}
