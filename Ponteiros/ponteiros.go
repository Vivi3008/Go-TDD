package ponteiros

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(value Bitcoin) {
	c.saldo += value
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
