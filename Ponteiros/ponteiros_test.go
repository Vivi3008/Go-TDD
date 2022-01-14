package ponteiros

import (
	"testing"
)

func TestPonteiros(t *testing.T) {
	t.Run("Teste Carteira Depositar", func(t *testing.T) {
		c := Carteira{}
		c.Depositar(50)
		expected := "50 BTC"
		verificaSaldo(t, c, expected)
	})
	t.Run("Teste carteira Retirar com saldo suficiente", func(t *testing.T) {
		cart := Carteira{saldo: 20}
		err := cart.Retirar(10)
		expected := "10 BTC"
		verificaSaldo(t, cart, expected)
		confirmaErro(t, nil, err)
	})
	t.Run("Retirar valor com saldo insuficiente", func(t *testing.T) {
		cart := Carteira{saldo: 20}
		err := cart.Retirar(100)
		verificaSaldo(t, cart, "20 BTC")
		confirmaErro(t, ErrInsufficientLimit, err)
	})
}

func verificaSaldo(t *testing.T, c Carteira, expected string) {
	t.Helper()
	got := c.GetSaldo().String()
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func confirmaErro(t *testing.T, wantErr, err error) {
	if err != wantErr {
		t.Errorf("Expected %s, got %s", wantErr, err)
	}
}
