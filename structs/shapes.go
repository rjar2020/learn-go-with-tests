package structs

import "math"

//Shape assit in decoupling Area() calculation from concrete shapes
type Shape interface {
	Area() float64
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
