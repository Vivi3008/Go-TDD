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

	t.Run("Atualiza uma definição existente", func(t *testing.T) {
		dic := Dicionario{"hello": "Hello World!"}
		errUpd := dic.Atualiza("hello", "Olá mundo!")

		comparaErro(t, nil, errUpd)

		got, err := dic.Busca("hello")
		expected := "Olá mundo!"

		comparaErro(t, nil, err)
		comparaStrings(t, expected, got)
	})

	t.Run("Atualiza uma palavra nova", func(t *testing.T) {
		err := dicionario.Atualiza("hello", "World")

		comparaErro(t, ErrAtualizaPalavra, err)
	})

	t.Run("Teste Deleta uma palavra", func(t *testing.T) {
		dic := Dicionario{"teste": "testando 1,2,3", "foo": "bar", "home": "sweet home"}
		dic.Deleta("home")

		_, err := dic.Busca("home")
		comparaErro(t, ErrPalavraInexistente, err)
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
