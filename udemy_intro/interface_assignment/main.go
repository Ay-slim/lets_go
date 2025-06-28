package main

import "fmt"

type triangle struct{
	height float64;
	base float64;
}
func (t triangle)getArea()float64{
	return 0.5 * t.height * t.base
}
type square struct{
	sideLength float64;
}
func (s square)getArea()float64{
	return s.sideLength * s.sideLength
}
type shape interface {
	printArea()
}
func (t triangle)printArea(){
	fmt.Println(t.getArea())
}
func (s square)printArea(){
	fmt.Println(s.getArea())
}

func main () {
	tri := triangle {
		height: 3,
		base: 4,
	}
	tri.printArea()
	sq := square {
		sideLength: 5,
	}
	sq.printArea()
}