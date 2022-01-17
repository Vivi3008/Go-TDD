package mapas

import (
	"testing"
)

func TestMapas(t *testing.T) {
	dicionario := Dicionario{"teste": "Isso é apenas um teste", "foo": "bar", "lar": "doce lar"}

	t.Run("Teste Mapas palavra conhecida", func(t *testing.T) {
		got, _ := dicionario.Busca("teste")
		expected := "Isso é apenas um teste"

		comparaStrings(t, expected, got)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("oi")
		if err == nil {
			t.Errorf("Expected err not be nil")
		}
	})
}

func comparaStrings(t *testing.T, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
