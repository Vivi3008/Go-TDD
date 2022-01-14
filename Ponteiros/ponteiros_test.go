package ponteiros

import (
	"testing"
)

func TestPonteiros(t *testing.T) {
	verificaSaldo := func(t *testing.T, c Carteira, expected string) {
		t.Helper()
		got := c.GetSaldo().String()
		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	}

	confirmaErro := func(t *testing.T, err error) {
		if err == nil {
			t.Errorf("Expected %s, got nil", err)
		}
	}
	t.Run("Teste Carteira Depositar", func(t *testing.T) {
		c := Carteira{}
		c.Depositar(50)
		expected := "50 BTC"
		verificaSaldo(t, c, expected)
	})
	t.Run("Teste carteira Retirar", func(t *testing.T) {
		cart := Carteira{saldo: 20}
		cart.Retirar(10)
		expected := "10 BTC"
		verificaSaldo(t, cart, expected)
	})
	t.Run("Retirar valor com saldo insuficiente", func(t *testing.T) {
		cart := Carteira{saldo: 20}
		err := cart.Retirar(100)
		verificaSaldo(t, cart, "20 BTC")
		confirmaErro(t, err)

	})
}
