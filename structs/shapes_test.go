package structs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			require.Equal(t, tt.hasArea, got, "The shape doesn't have the expected area")
		})
	}
}

func TestPerimeter(t *testing.T) {
	areaTests := []struct {
		name         string
		shape        PolygonalShape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPerimeter: 36},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasPerimeter: 28.97056274847714},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			require.Equal(t, tt.hasPerimeter, got, "The shape doesn't have the expected area")
		})
	}
}

func TestCircumference(t *testing.T) {
	areaTests := []struct {
		name             string
		shape            EllipseShape
		hasCircumference float64
	}{
		{name: "Circle", shape: Circle{Radius: 10}, hasCircumference: 62.83185307179586},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Circumference()
			require.Equal(t, tt.hasCircumference, got, "The shape doesn't have the expected area")
		})
	}
}

func BenchmarkArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaTests := []struct {
			name    string
			shape   Shape
			hasArea float64
		}{
			{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
			{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
			{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
		}
		for _, tt := range areaTests {
			tt.shape.Area()
		}
	}
}

func BenchmarkPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaTests := []struct {
			name         string
			shape        PolygonalShape
			hasPerimeter float64
		}{
			{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPerimeter: 36},
			{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasPerimeter: 28.97056274847714},
		}
		for _, tt := range areaTests {
			tt.shape.Perimeter()
		}
	}
}

func BenchmarkCircumference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		areaTests := []struct {
			name             string
			shape            EllipseShape
			hasCircumference float64
		}{
			{name: "Circle", shape: Circle{Radius: 10}, hasCircumference: 62.83185307179586},
		}
		for _, tt := range areaTests {
			tt.shape.Circumference()
		}
	}
}
func ExampleTriangle_Area() {
	triangle := Triangle{Base: 12, Height: 6}
	fmt.Print(triangle.Area())
	/* Output: 36*/
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{Width: 12, Height: 6}
	fmt.Print(rectangle.Area())
	/* Output: 72*/
}

func ExampleCircle_Area() {
	circle := Circle{Radius: 10}
	fmt.Print(circle.Area())
	/* Output: 314.1592653589793*/
}

func ExampleRectangle_Perimeter() {
	rectangle := Rectangle{Width: 12, Height: 6}
	fmt.Print(rectangle.Perimeter())
	/* Output: 36*/
}

func ExampleTriangle_Perimeter() {
	triangle := Triangle{Base: 12, Height: 6}
	fmt.Print(triangle.Perimeter())
	/* Output: 28.97056274847714*/
}

func ExampleCircle_Circumference() {
	circle := Circle{Radius: 10}
	fmt.Print(circle.Circumference())
	/* Output: 62.83185307179586*/
}
