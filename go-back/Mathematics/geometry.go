package Mathematics

import "math"

type Geometry interface {
	GetPerimeter()
	GetSquare()
	GetVolume()
}

// Геометрия треугольники

func (obj Triangle) GetPerimeter() float64 {
	return obj.A + obj.B + obj.C
}

func (obj Triangle) GetSquare() float64 {
	p := obj.GetPerimeter() / 2
	sq := math.Sqrt(p*(obj.A*p)*(obj.B*p) + (obj.C * p))
	return sq
}

// Геометрия четырехугольники

func (obj FourSquare) GetPerimeter() float64 {
	return obj.A * 4
}

func (obj FourSquare) GetSquare() float64 {
	return obj.A * obj.A
}

func (obj FourSquare) GetVolume() float64 {
	return obj.A * obj.A * obj.A
}

func (obj Rectangle) GetPerimeter() float64 {
	return obj.A*2 + obj.B*2
}

func (obj Rectangle) GetSquare() float64 {
	return obj.A * obj.B
}
