package sync

type Contador struct {
	valor int
}

func (c *Contador) Incrementa() {
	c.valor++
}

func (v *Contador) Valor() int {
	return v.valor
}
