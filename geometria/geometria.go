package geometria

import "math"

type Geometria interface{
	Area() float32
	Perimeter() float32
}

type Square struct{
	Height float32
	Width float32
}

type Circle struct {
	Rad float32
}

func NewSquare(height, width float32) Geometria{
	return Square{
		Height: height,
		Width: width,
	}
}

func NewCircle(rad float32) Geometria{
	return Circle{
		Rad: rad,
	}
}

func (sqr Square) Area()  float32 {
	return sqr.Height * sqr.Width
}

func (sqr Square) Perimeter()  float32 {
	return 2 * sqr.Height + 2 * sqr.Width
}

func (c Circle) Area() float32 {
	return  math.Pi * c.Rad * c.Rad
}

func (c Circle) Perimeter() float32 {
	return  2 * math.Pi * c.Rad
}