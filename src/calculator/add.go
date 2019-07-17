package main

import (
  "fmt"
)

func add(x, y int) int {
    return x + y
}

func main () {
    var x,y = 5, 5

    total := add(x, y)
    fmt.Printf("Total da Soma de %v + %v é %v", x, y, total);
}
