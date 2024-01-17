package basic

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func structure() {
	var v Vertex
	fmt.Println(v)

	v = Vertex{}
	fmt.Println(v)

	v = Vertex{1.0, 2.0}
	fmt.Println(v)

	v = Vertex{Y: 2.5}
	fmt.Println(v)

	v = Vertex{X: 1.0, Y: 3.0}
	fmt.Println(v)

	v.X *= 1.5
	v.Y *= 2.5

	fmt.Println(v)

	var vPointer = &Vertex{1, 3}
	fmt.Println(vPointer)

	vPointer.X, vPointer.Y = vPointer.Y, vPointer.X
	fmt.Println(vPointer)
}

func (v *Vertex) Square() {
	v.X *= v.X
	v.Y *= v.Y
}

type MyMap map[string]int

func (m MyMap) Length() int {
	return len(m)
}

func methods() {
	vert := Vertex{3, 4}
	fmt.Println(vert)

	vert.Square()
	fmt.Println(vert)

	mm := MyMap{"A": 1, "B": 2}

	fmt.Println(mm)
	fmt.Println(mm["A"])
	fmt.Println(mm.Length())
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("%T's area is %0.2f\n", s, s.Area())
}

func interfaces() {
	r := Rectangle{
		Width:  5,
		Height: 10,
	}
	PrintArea(r)

	c := Circle{
		Radius: 5,
	}
	PrintArea(c)

	var s Shape = Circle{}

	c2 := s.(Circle)
	fmt.Printf("%T\n", c2)
}
