package main

import (
  "flag"
  "fmt"
)


func main() {
  var addr = flag.String("addr", ":8080", "http service address")
  var species = flag.String("species", "gopher", "the species we are studying")

  flag.Parse()

  fmt.Printf("Result: %v\n", *addr)
  fmt.Printf("Result: %v\n", *species)
}
