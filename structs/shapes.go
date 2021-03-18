package structs

import "math"

//Shape assist in decoupling Area() calculation from concrete shapes
type Shape interface {
	Area() float64
}

//EllipseShape assist in decoupling Circumference() calculation for ellipses
type EllipseShape interface {
	Shape
	Circumference() float64
}

//PolygonalShape assist in decoupling Perimeter() calculation for Polygons
type PolygonalShape interface {
	Shape
	Perimeter() float64
}

//Rectangle models a rectangular shape
type Rectangle struct {
	Width  float64
	Height float64
}

//Area calculates a rectangle' area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Perimeter calculates a rectangle' perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

//Circle models a round shape
type Circle struct {
	Radius float64
}

//Area calculates a circle' area
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Circumference calculates a circle' perimeter
func (c Circle) Circumference() float64 {
	return math.Pi * c.Radius * 2
}

//Triangle models a triangular shape
type Triangle struct {
	Base   float64
	Height float64
}

//Area calculates a triangle' area
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

//Perimeter calculates a triangle' perimeter, based on base and height, so an approximation
func (t Triangle) Perimeter() float64 {
	return t.Base + math.Sqrt(math.Pow(t.Base, 2)+(math.Pow(t.Height, 2)*4))
}
