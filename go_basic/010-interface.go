package gobasic

import "fmt"

// Shape
type Shape interface {
	area() float64
}

// define method of Shape
func getArea(s Shape) float64 {
	return s.area()
}

// Cicle
type Cicle struct {
	radius float64
	x      float64
	y      float64
}

// define method of Cicle
func (c Cicle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// Rectangle
type Rectangle struct {
	width  float64
	height float64
}

// define method of Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func ShowInterface() {
	circle := Cicle{radius: 5, x: 0, y: 0}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Println("Circle area:", getArea(circle))
	fmt.Println("Rectangle area:", getArea(rectangle))

}
