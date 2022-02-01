package reflexao

import "testing"

func TestReflection(t *testing.T) {
	casos := []struct {
		Name string
		args interface{}
		want []string
	}{
		{
			Name: "Struct com um campo string",
			args: struct {
				Nome string
			}{"Chris"},
			want: []string{"Chris"},
		},
	}

	for _, tc := range casos {
		tt := tc //captura a variavel tc
		t.Run(tt.Name, func(t *testing.T) {
			var result []string

			Percorre(tt.args, func(entrada string) {
				result = append(result, entrada)
			})
			if result[0] != tt.want[0] {
				t.Errorf("expected %s, got %s", tt.want[0], result[0])
			}
		})

	}
}
