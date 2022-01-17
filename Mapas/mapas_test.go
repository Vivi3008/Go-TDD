package mapas

import "testing"

func TestMapas(t *testing.T) {
	t.Run("Teste Mapas", func(t *testing.T) {
		dicionario := map[string]string{"Teste": "Isso é apenas um teste"}
		got := Busca(dicionario, "Teste")
		expected := "Isso é apenas um teste"

		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}
