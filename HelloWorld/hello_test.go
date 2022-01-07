package helloworld

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	verifyMessageTest := func(t *testing.T, expected string, got string) {
		t.Helper()

		if expected != got {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		result := Hello("Jane")
		expected := "Hello Jane"

		verifyMessageTest(t, expected, result)
	})

	t.Run("Say hello world if has empty string", func(t *testing.T) {
		got := Hello("")
		expected := "Hello World"

		verifyMessageTest(t, expected, got)
	})
}
