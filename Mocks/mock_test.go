package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestContagem(t *testing.T) {
	t.Run("Teste contagem", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SleeperSpy{}

		Contagem(buffer, sleeperSpy)

		result := buffer.String()
		expected := `3
2
1
vai!`
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
	t.Run("pausa antes de cada impressao", func(t *testing.T) {
		spy := &SpyContagemOperacoes{}
		Contagem(spy, spy)

		expected := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(expected, spy.Chamadas) {
			t.Errorf("Expected %s, got %s", expected, spy.Chamadas)
		}
	})
}
