package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestContagem(t *testing.T) {
	t.Run("Teste contagem", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SleeperSpy{}

		Contagem(buffer, sleeperSpy)
		fmt.Println(sleeperSpy)
		result := buffer.String()
		expected := `3
2
1
vai!`

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}