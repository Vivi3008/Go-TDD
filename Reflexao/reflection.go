package reflexao

import "reflect"

func Percorre(x interface{}, fn func(entrada string)) {
	valor := reflect.ValueOf(x)
	campo := valor.Field(0)

	fn(campo.String())
}
