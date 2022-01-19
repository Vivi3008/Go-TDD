package main

import (
	"bytes"
	"testing"
)

func TestContagem(t *testing.T) {
	t.Run("Teste contagem", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Contagem(buffer)

		result := buffer.String()
		expected := "3"

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
