package structs

import "math"

type Shape interface {
	Area() float64
}

//Rectangle models a rectangular shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle models a round shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Perimeter calculates a rectangles' perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area calculates a rectangles' area
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
