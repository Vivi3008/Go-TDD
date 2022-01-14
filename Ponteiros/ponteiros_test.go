package ponteiros

import (
	"testing"
)

func TestPonteiros(t *testing.T) {
	t.Run("Teste Carteira", func(t *testing.T) {
		c := Carteira{}
		c.Depositar(50)
		got := c.Saldo().String()
		expected := "50 BTC"

		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}
