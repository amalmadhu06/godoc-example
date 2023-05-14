package area

// Square represents a square shape.
type Square struct {
	Side float64 // Length of the side of the square
}

// Rectangle represents a rectangle shape.
type Rectangle struct {
	Length float64 // Length of the rectangle
	Width  float64 // Width of the rectangle
}

// Circle represents a circle shape.
type Circle struct {
	Radius float64 // Radius of the circle
}

// Triangle represents a triangle shape.
type Triangle struct {
	Base   float64 // Length of the base of the triangle
	Height float64 // Height of the triangle
}

// Trapezoid represents a trapezoid shape.
type Trapezoid struct {
	Base1  float64 // Length of the first base of the trapezoid
	Base2  float64 // Length of the second base of the trapezoid
	Height float64 // Height of the trapezoid
}

// Parallelogram represents a parallelogram shape.
type Parallelogram struct {
	Base   float64 // Length of the base of the parallelogram
	Height float64 // Height of the parallelogram
}

// Rhombus represents a rhombus shape.
type Rhombus struct {
	Diagonal1 float64 // Length of the first diagonal of the rhombus
	Diagonal2 float64 // Length of the second diagonal of the rhombus
}

// RegularPentagon represents a regular pentagon shape.
type RegularPentagon struct {
	Side float64 // Length of the side of the regular pentagon
}

// RegularHexagon represents a regular hexagon shape.
type RegularHexagon struct {
	Side float64 // Length of the side of the regular hexagon
}
