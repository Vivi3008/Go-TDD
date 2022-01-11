package iteracao

import "testing"

func TestIteracao(t *testing.T) {
	t.Run("Shoul repeat a caracter 5 times", func(t *testing.T) {
		got := Repeat("a")
		expected := "aaaaa"

		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}
