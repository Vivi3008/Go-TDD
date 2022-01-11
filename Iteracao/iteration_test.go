package iteracao

import (
	"fmt"
	"testing"
)

func TestIteracao(t *testing.T) {
	t.Run("Shoul repeat a caracter 5 times", func(t *testing.T) {
		got := Repeat("a", 6)
		expected := "aaaaaa"

		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("vi", 2)
	fmt.Println(res)
	// Output vivi
}
