package sync

import "testing"

func TestSync(t *testing.T) {
	t.Run("Teste contador", func(t *testing.T) {
		contador := Contador{}

		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		verificaContador(t, 3, contador)
	})
}

func verificaContador(t *testing.T, expected int, contador Contador) {
	t.Helper()

	if contador.Valor() != expected {
		t.Errorf("Expected contador have %v, but have %v", expected, contador.Valor())
	}
}
