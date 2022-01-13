package interfaces

import "testing"

func TestPerimetro(t *testing.T) {
	t.Run("Calcula perimetro", func(t *testing.T) {
		ret := Retangulo{
			Larg: 10.0,
			Alt:  10.0,
		}
		got := ret.Perimetro()
		expected := 40.0

		if got != expected {
			t.Errorf("Expected %.2f, got %.2f", expected, got)
		}
	})
}

func TestArea(t *testing.T) {
	testCases := []struct {
		name string
		arg  Forma
		want float64
	}{
		{
			name: "Calcula area retangulo",
			arg:  Retangulo{10, 10},
			want: 100.0,
		},
		{
			name: "Calcula area do circulo",
			arg:  Circulo{10},
			want: 314.1592653589793,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			got := tt.arg.Area()
			if got != tt.want {
				t.Errorf("Expected %.2f, got %.2f", tt.want, got)
			}
		})
	}
}
