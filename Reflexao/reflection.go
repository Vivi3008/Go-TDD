package reflexao

import "reflect"

func Percorre(x interface{}, fn func(entrada string)) {
	valor := reflect.ValueOf(x)

	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)

		if campo.Kind() == reflect.String {
			fn(campo.String())
		}

		if campo.Kind() == reflect.Struct {
			Percorre(campo.Interface(), fn)
		}
	}
}
