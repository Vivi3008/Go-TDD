package generics

import (
	"testing"
)

func Test_generics(t *testing.T) {
	tests := []struct {
		name      string
		argsInt   []int
		argsFloat []float64
		wantInt   int
		wantFloat float64
	}{
		{
			name:    "Should sum int values",
			argsInt: []int{6, 5},
			wantInt: 11,
		},
		{
			name:      "Should sum float values",
			argsFloat: []float64{6.4, 5.2},
			wantFloat: 12,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			if tt.argsFloat != nil {
				got := sumIntsOrFloas(tt.argsInt)

				if tt.wantInt != got {
					t.Errorf("Expected %d, got %d", tt.wantInt, got)
				}
			} else {
				got := sumIntsOrFloas(tt.argsFloat)

				if tt.wantFloat != got {
					t.Errorf("Expected %.2f, got %.2f", tt.wantFloat, got)
				}
			}
		})
	}
}
