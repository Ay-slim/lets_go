package main

import "fmt"
import "math"

// Structs, think of these as objects or classes with properties and methods
type Animals struct {
	name string
	sound string
}

// Struct methods (function property of a struct)
func (animal Animals) makeSound() string {
	fmt.Println(animal.sound + " " + animal.sound + " " + animal.sound)
	return animal.name
}

// Interfaces: An interface is a typed collection of methods
type geometry interface { // This defines any type that has the area and parameter methods
	area() float64
	perimeter() float64
}

type rect struct {
	length float64
	width float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure (g geometry) {
	fmt.Println(g, "measurement base")
	fmt.Println(g.area(), "measurement area")
	fmt.Println(g.perimeter(), "measurement perimeter")
}

func main() {
	doggie := Animals{"dog", "bark"}
	fmt.Println(doggie.name)
	fmt.Println(doggie.sound)
	doggie.makeSound() // Should print bark 3 times
	fmt.Println(doggie.makeSound()) // Should print "bark bark bark" and "dog"

	rectangle := rect{2, 3}
	fmt.Println(rectangle.area(), "Rectangle Area, should be 6")
	fmt.Println(rectangle.perimeter(), "Rectangle Perimeter, should be 10")
	circ := circle{5}
	fmt.Println(circ.area(), "Circle Area, should be 25 pi, between 75 and 80")
	fmt.Println(circ.perimeter(), "Circle Perimeter, should be 10 pi; 31.4 ish")
	measure(rectangle)
	measure(circ)
}
