// Package challenge10 contains the solution for Challenge 10.
package challenge10

import (
	"fmt"
	"math"
	"slices"
	// Add any necessary imports here
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
	fmt.Stringer // Includes String() string method
}

// Rectangle represents a four-sided shape with perpendicular sides
type Rectangle struct {
	Width  float64
	Height float64
}

// NewRectangle creates a new Rectangle with validation
func NewRectangle(width, height float64) (*Rectangle, error) {
	// TODO: Implement validation and construction
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("can't create rectangle with height %f, width %f", height, width)
	}

	return &Rectangle{
		Width:  width,
		Height: height,
	}, nil
}

// Area calculates the area of the rectangle
func (r *Rectangle) Area() float64 {
	// TODO: Implement area calculation
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	return r.Height*2 + r.Width*2
	// TODO: Implement perimeter calculation
}

// String returns a string representation of the rectangle
func (r *Rectangle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Shape: %s, Height: %f, Width: %f", "Rectangle", r.Height, r.Width)
}

// Circle represents a perfectly round shape
type Circle struct {
	Radius float64
}

// NewCircle creates a new Circle with validation
func NewCircle(radius float64) (*Circle, error) {
	// TODO: Implement validation and construction

	if radius <= 0 {
		return nil, fmt.Errorf("radius must be positive")
	}
	return &Circle{
		radius,
	}, nil
}

// Area calculates the area of the circle
func (c *Circle) Area() float64 {
	// TODO: Implement area calculation
	return math.Pow(c.Radius, 2) * math.Pi
}

// Perimeter calculates the circumference of the circle
func (c *Circle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return 2 * c.Radius * math.Pi
}

// String returns a string representation of the circle
func (c *Circle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Shape: %s. Radius: %f", "Circle", c.Radius)
}

// Triangle represents a three-sided polygon
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

// NewTriangle creates a new Triangle with validation
func NewTriangle(a, b, c float64) (*Triangle, error) {
	// TODO: Implement validation and construction
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, fmt.Errorf("can't create a triangle with side smaller or equal than 0")
	}
	if a+b <= c || b+c <= a || a+c <= b {
		return nil, fmt.Errorf("can't create a triangle with one side smaller than the sum of the other 2")
	}
	return &Triangle{a, b, c}, nil
}

// Area calculates the area of the triangle using Heron's formula
func (t *Triangle) Area() float64 {
	// TODO: Implement area calculation using Heron's formula
	s := 0.5 * t.Perimeter()
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

// Perimeter calculates the perimeter of the triangle
func (t *Triangle) Perimeter() float64 {
	// TODO: Implement perimeter calculation
	return t.SideA + t.SideB + t.SideC
}

// String returns a string representation of the triangle
func (t *Triangle) String() string {
	// TODO: Implement string representation
	return fmt.Sprintf("Shape: Triangle, sides: %f, %f, %f", t.SideA, t.SideB, t.SideC)
}

// ShapeCalculator provides utility functions for shapes
type ShapeCalculator struct{}

// NewShapeCalculator creates a new ShapeCalculator
func NewShapeCalculator() *ShapeCalculator {
	// TODO: Implement constructor
	return &ShapeCalculator{}
}

// PrintProperties prints the properties of a shape
func (sc *ShapeCalculator) PrintProperties(s Shape) {
	// TODO: Implement printing shape properties
	fmt.Print(s.String())
}

// TotalArea calculates the sum of areas of all shapes
func (sc *ShapeCalculator) TotalArea(shapes []Shape) float64 {
	// TODO: Implement total area calculation
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// LargestShape finds the shape with the largest area
func (sc *ShapeCalculator) LargestShape(shapes []Shape) Shape {
	// TODO: Implement finding largest shape
	largest := shapes[0]
	for _, s := range shapes {
		if largest.Area() < s.Area() {
			largest = s
		}
	}
	return largest
}

// SortByArea sorts shapes by area in ascending or descending order
func (sc *ShapeCalculator) SortByArea(shapes []Shape, ascending bool) []Shape {
	// TODO: Implement sorting shapes by area
	var comp func(Shape, Shape) int
	if ascending {
		comp = func(s1 Shape, s2 Shape) int {
			return int(math.Ceil(s1.Area() - s2.Area()))
		}
	} else {
		comp = func(s1 Shape, s2 Shape) int {
			return int(math.Ceil(s2.Area() - s1.Area()))
		}
	}
	slices.SortFunc(shapes, comp)
	return shapes
}
