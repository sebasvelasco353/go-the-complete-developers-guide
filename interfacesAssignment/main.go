package main

import "fmt"

type shape interface {
  getArea() float64
}

type square struct {
  sideLength float64
}

type triangle struct {
  base float64
  height float64
}

func main() {
  tr := triangle{
    base: 3.8,
    height: 12.4,
  }
  sq := square{
    sideLength: 23.5,
  }

  printArea(tr)
  printArea(sq)
}

func (sq square)getArea()float64 {
  area := sq.sideLength * sq.sideLength
  return area
}

func (tr triangle)getArea()float64 {
  area := 0.5 * tr.base * tr.height
  return area
}

func printArea(s shape) {
  fmt.Println(s.getArea())
}

