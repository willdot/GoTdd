package structs

import "math"

// Shape is an interface for shapes
type Shape interface {
	Area() float64
}

// Rectangle containing width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle containing radius
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle containing base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter returns the permiter of a rectangle given the width and height
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Area returns the area of a rectangle given the width and height
func Area(width float64, height float64) float64 {
	return width * height
}

// PerimeterStruct returns the permiter of a rectangle given the width and height from a rectangle struct
func PerimeterStruct(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// AreaStruct returns the area of a rectangle given the width and height from a rectangle struct
func AreaStruct(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
