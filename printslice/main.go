package main

import "fmt"

func main() {
  numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  for _, number := range numbers {
    if number % 2 == 0 {
      fmt.Printf("The number %v is even\n", number)
    } else {
      fmt.Printf("the number %v is odd\n", number)
    }
  }
}
