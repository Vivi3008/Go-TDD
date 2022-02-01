package reflexao

import "reflect"

func Percorre(x interface{}, fn func(entrada string)) {
	valor := extractValue(x)

	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)

		switch campo.Kind() {
		case reflect.String:
			fn(campo.String())
		case reflect.Struct:
			Percorre(campo.Interface(), fn)
		}
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
