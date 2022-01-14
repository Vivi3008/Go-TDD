package ponteiros

import "fmt"

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}
type Stringer interface {
	String() string
}

func (c *Carteira) Depositar(value Bitcoin) {
	c.saldo += value
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
