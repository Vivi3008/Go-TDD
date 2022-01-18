package main

import (
	"bytes"
	"testing"
)

func TestInjection(t *testing.T) {
	t.Run("Teste cumprimentar", func(t *testing.T) {
		buffer := bytes.Buffer{}
		SayHello(&buffer, "viviane")

		result := buffer.String()
		expected := "Olá viviane"

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
