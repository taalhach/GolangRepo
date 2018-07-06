package main

import "fmt"

type Point struct {
	X,Y int
}
type Wheel struct {
	Point
	radius int
}

func main() {
	rw:=Wheel{
		Point{1,2},
		15,
	}
	fmt.Println(rw)
	rw.X=10
	rw.Y=15           // equal to rw.pointer to.x
	fmt.Println(rw)

}
