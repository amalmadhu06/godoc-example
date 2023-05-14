// Package area provides functions for calculating the area of various shapes.
//
// This package includes functions to calculate the area of the following shapes:
//   - Square
//   - Rectangle
//   - Circle
//   - Triangle
//   - Trapezoid
//   - Parallelogram
//   - Rhombus
//   - Regular Pentagon
//   - Regular Hexagon
package area

import "math"

// Pi represents the value of Pi
const Pi = 3.14

// Area calculates the area of a square.
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Area calculates the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Area calculates the area of a circle.
func (c Circle) Area() float64 {
	return Pi * c.Radius * c.Radius
}

// Area calculates the area of a triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Area calculates the area of a trapezoid.
func (t Trapezoid) Area() float64 {
	return ((t.Base1 + t.Base2) * t.Height) / 2
}

// Area calculates the area of a parallelogram.
func (p Parallelogram) Area() float64 {
	return p.Base * p.Height
}

// Area calculates the area of a rhombus.
func (r Rhombus) Area() float64 {
	return (r.Diagonal1 * r.Diagonal2) / 2
}

// Area calculates the area of a regular pentagon.
func (p RegularPentagon) Area() float64 {
	return (p.Side * p.Side * math.Sqrt(5*(5+2*math.Sqrt(5)))) / 4
}

// Area calculates the area of a regular hexagon.
func (h RegularHexagon) Area() float64 {
	return (3 * math.Sqrt(3) * h.Side * h.Side) / 2
}
