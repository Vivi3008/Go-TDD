package arrayseslices

import "testing"

func TestArrayESlice(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Soma(numbers)
	expected := 15

	if got != expected {
		t.Errorf("Expected %d, got %d, data %v", expected, got, numbers)
	}
}
