package ponteiros

import (
	"errors"
	"fmt"
)

var ErrInsufficientLimit = errors.New("saldo insuficiente")

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

func (c *Carteira) Retirar(value Bitcoin) error {
	if value > c.saldo {
		return ErrInsufficientLimit
	}
	c.saldo -= value
	return nil
}

func (c *Carteira) GetSaldo() Bitcoin {
	return c.saldo
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
