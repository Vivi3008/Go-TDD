package mapas

import (
	"testing"
)

func TestMapas(t *testing.T) {
	dicionario := Dicionario{"teste": "Isso é apenas um teste", "foo": "bar", "lar": "doce lar"}

	t.Run("Teste Mapas palavra conhecida", func(t *testing.T) {
		got, err := dicionario.Busca("teste")
		expected := "Isso é apenas um teste"

		comparaStrings(t, expected, got)
		comparaErro(t, nil, err)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("oi")
		comparaErro(t, ErrPalavraInexistente, err)
	})

	t.Run("adiciona nova palavra", func(t *testing.T) {
		word := Dicionario{}
		word.Adiciona("oi", "vou responder Oi!")
		got, err := word.Busca("oi")
		expected := "vou responder Oi!"
		comparaErro(t, nil, err)
		comparaStrings(t, expected, got)
	})

	t.Run("nao adiciona palavra ja existente", func(t *testing.T) {
		errAdd := dicionario.Adiciona("lar", "quero mudar definicao")
		comparaErro(t, ErrPalavraJaExiste, errAdd)
		got, err := dicionario.Busca("lar")
		comparaErro(t, nil, err)
		comparaStrings(t, got, "doce lar")
	})

}

func comparaStrings(t *testing.T, expected, got string) {
	t.Helper()
	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func comparaErro(t *testing.T, wantErr, err error) {
	t.Helper()
	if wantErr != err {
		t.Errorf("Expected %s, got %s", wantErr, err)
	}
}
