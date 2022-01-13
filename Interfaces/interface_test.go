package interfaces

import "testing"

func TestPerimetro(t *testing.T) {
	t.Run("Calcula perimetro", func(t *testing.T) {
		ret := Retangulo{
			Larg: 10.0,
			Alt:  10.0,
		}
		got := ret.Perimetro()
		expected := 40.0

		if got != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, got)
		}
	})
}

func TestArea(t *testing.T) {
	calculaArea := func(t *testing.T, forma Forma, expected float64) {
		t.Helper()
		got := forma.Area()
		if got != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, got)
		}
	}

	t.Run("Calcula Area retangulo", func(t *testing.T) {
		ret := Retangulo{
			Larg: 10.0,
			Alt:  10.0,
		}
		calculaArea(t, ret, 100.0)
	})
	t.Run("Calcula area circulo", func(t *testing.T) {
		circ := Circulo{10}
		expected := 314.1592653589793
		calculaArea(t, circ, expected)
	})
}
