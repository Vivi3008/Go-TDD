package sync

import "testing"

func TestSync(t *testing.T) {
	t.Run("Teste contador", func(t *testing.T) {
		contador := Contador{}

		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		if contador.Valor() != 3 {
			t.Errorf("Expected contador have %v, but have %v", 3, contador.Valor())
		}
	})
}
