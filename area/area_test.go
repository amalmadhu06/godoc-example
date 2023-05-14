package area_test

import (
	"fmt"
	"github.com/amalmadhu06/godoc-example/area"
)

func Example() {
	c := area.Circle{Radius: 10}
	fmt.Println(c.Area())
	// Output:
	// 314
}

func ExampleSquare_Area() {
	s := area.Square{Side: 10}
	fmt.Println(s.Area())
	// Output:
	// 100
}

func ExampleRectangle_Area() {
	r := area.Rectangle{
		Length: 15,
		Width:  10,
	}
	fmt.Println(r.Area())
	// Output:
	// 150
}

func ExampleCircle_Area() {
	c := area.Circle{Radius: 10}
	fmt.Println(c.Area())
	// Output:
	// 314
}

func ExampleTriangle_Area() {
	t := area.Triangle{
		Base:   10,
		Height: 5,
	}
	fmt.Println(t.Area())
	// Output:
	// 25
}

func ExampleTrapezoid_Area() {
	t := area.Trapezoid{
		Base1:  12,
		Base2:  18,
		Height: 7,
	}
	fmt.Println(t.Area())
	// Output:
	// 105
}

func ExampleParallelogram_Area() {
	p := area.Parallelogram{
		Base:   15,
		Height: 5,
	}
	fmt.Println(p.Area())
	// Output:
	// 75
}

func ExampleRhombus_Area() {
	r := area.Rhombus{
		Diagonal1: 10,
		Diagonal2: 12,
	}
	fmt.Println(r.Area())
	// Output:
	// 60
}

func ExampleRegularPentagon_Area() {
	rp := area.RegularPentagon{Side: 10}
	fmt.Println(rp.Area())
	// Output:
	// 172.0477400588967
}

func ExampleRegularHexagon_Area() {
	rh := area.RegularHexagon{Side: 10}
	fmt.Println(rh.Area())
	// Output:
	// 259.8076211353316
}
