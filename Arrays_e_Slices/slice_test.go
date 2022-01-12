package arrayseslices

import (
	"reflect"
	"testing"
)

func TestArrayESlice(t *testing.T) {
	verificaResultado := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	t.Run("Array de cinco numeros", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Soma(numbers)
		expected := 15

		if got != expected {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})
	t.Run("Funcao soma tudo", func(t *testing.T) {
		got := SomaTodoResto([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		verificaResultado(t, got, expected)
	})
	t.Run("Soma slices vazios", func(t *testing.T) {
		got := SomaTodoResto([]int{}, []int{1, 3, 5})
		expected := []int{0, 8}

		verificaResultado(t, got, expected)
	})
}
