package inteiros

import (
	"fmt"
	"testing"
)

func TestInteiros(t *testing.T) {
	t.Run("Should sum two values", func(t *testing.T) {
		got := Sum(2, 1)
		expected := 3

		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})
}

func ExampleSum() {
	res := Sum(3, 3)
	fmt.Println(res)
	// Output 6
}
