package reflexao

import "reflect"

func Percorre(x interface{}, fn func(entrada string)) {
	valor := reflect.ValueOf(x)

	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)
		fn(campo.String())
	}
}
