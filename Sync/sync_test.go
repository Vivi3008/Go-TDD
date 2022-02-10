package sync

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	t.Run("Teste contador", func(t *testing.T) {
		contador := NewContador()

		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		verificaContador(t, 3, contador)
	})

	t.Run("roda concorrentemente em segurança", func(t *testing.T) {
		contagemEsperada := 1000
		contador := NewContador()

		var wg sync.WaitGroup
		wg.Add(contagemEsperada)

		for i := 0; i < contagemEsperada; i++ {
			go func(w *sync.WaitGroup) {
				contador.Incrementa()
				w.Done()
			}(&wg)
		}

		wg.Wait()
		verificaContador(t, contagemEsperada, contador)
	})
}

func verificaContador(t *testing.T, expected int, contador *Contador) {
	t.Helper()

	if contador.Valor() != expected {
		t.Errorf("Expected contador have %v, but have %v", expected, contador.Valor())
	}
}
