package reflexao

import "reflect"

func Percorre(x interface{}, fn func(entrada string)) {
	valor := extractValue(x)

	qtdValores := 0
	var obtemCampo func(int) reflect.Value

	switch valor.Kind() {
	case reflect.String:
		fn(valor.String())
	case reflect.Slice:
		qtdValores = valor.Len()
		obtemCampo = valor.Index
	case reflect.Struct:
		qtdValores = valor.NumField()
		obtemCampo = valor.Field
	}

	for i := 0; i < qtdValores; i++ {
		Percorre(obtemCampo(i).Interface(), fn)
	}
}

func extractValue(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}
	return valor
}

//https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/reflection#escreva-o-teste-primeiro-6
