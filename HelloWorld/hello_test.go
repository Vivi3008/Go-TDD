package helloworld

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("Test Hello World", func(t *testing.T) {
		result := Hello("Jane")
		expected := "Hello Jane"

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
