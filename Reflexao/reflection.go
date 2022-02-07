package reflexao

import "reflect"

func Percorre(x interface{}, fn func(entrada string)) {
	valor := extractValue(x)

	percorreValor := func(valor reflect.Value) {
		Percorre(valor.Interface(), fn)
	}

	switch valor.Kind() {
	case reflect.String:
		fn(valor.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < valor.Len(); i++ {
			percorreValor(valor.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < valor.NumField(); i++ {
			percorreValor(valor.Field(i))
		}
	case reflect.Map:
		for _, key := range valor.MapKeys() {
			percorreValor(valor.MapIndex(key))
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
