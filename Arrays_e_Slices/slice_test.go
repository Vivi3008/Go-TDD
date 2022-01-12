package arrayseslices

import (
	"reflect"
	"testing"
)

func TestArrayESlice(t *testing.T) {
	t.Run("Array de cinco numeros", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Soma(numbers)
		expected := 15

		if got != expected {
			t.Errorf("Expected %d, got %d, data %v", expected, got, numbers)
		}
	})
	t.Run("Funcao soma tudo", func(t *testing.T) {
		got := SomaTudo([]int{1, 3, 2}, []int{1, 4})
		expected := []int{6, 5}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}
