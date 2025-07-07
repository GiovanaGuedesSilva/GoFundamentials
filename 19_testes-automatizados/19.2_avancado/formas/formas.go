package formas

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Base   float64
	Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Base * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Raio * c.Raio
}
