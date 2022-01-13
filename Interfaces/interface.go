package interfaces

import "math"

type Retangulo struct {
	Larg float64
	Alt  float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

type Forma interface {
	Area() float64
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Larg + r.Alt)
}

func (r Retangulo) Area() float64 {
	return r.Alt * r.Larg
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func (t Triangulo) Area() float64 {
	return (t.Altura * t.Base) / 2
}
