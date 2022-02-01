package reflexao

import (
	"reflect"
	"testing"
)

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
		{
			Name: "Struct com dois campos string",
			args: struct {
				Nome      string
				Sobrenome string
			}{"Chris", "Oliver"},
			want: []string{"Chris", "Oliver"},
		},
	}

	for _, tc := range casos {
		tt := tc //captura a variavel tc
		t.Run(tt.Name, func(t *testing.T) {
			var result []string

			Percorre(tt.args, func(entrada string) {
				result = append(result, entrada)
			})

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, result)
			}
		})

	}
}
