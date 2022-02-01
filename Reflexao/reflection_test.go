package reflexao

import (
	"reflect"
	"testing"
)

type Pessoa struct {
	Nome string
	Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}

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
		{
			Name: "Struct sem campo tipo string",
			args: struct {
				Nome  string
				Idade int
			}{"Chris", 36},
			want: []string{"Chris"},
		},
		{
			Name: "Struct aninhada",
			args: Pessoa{
				"Chris",
				Perfil{36, "Goiânia"},
			},
			want: []string{"Chris", "Goiânia"},
		},
		{
			Name: "Ponteiros para coisas",
			args: &Pessoa{
				"Chris",
				Perfil{36, "Goiânia"},
			},
			want: []string{"Chris", "Goiânia"},
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
