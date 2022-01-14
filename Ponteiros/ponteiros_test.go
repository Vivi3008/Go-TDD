package ponteiros

import (
	"testing"
)

func TestPonteiros(t *testing.T) {
	t.Run("Teste Carteira", func(t *testing.T) {
		c := Carteira{}
		c.Depositar(50)
		got := c.Saldo()
		expected := Bitcoin(50)

		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})
}
