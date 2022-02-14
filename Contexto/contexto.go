package contexto

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Println(store.Fetch())
	}
}

//https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/contexto#escreva-codigo-o-suficiente-para-fazer-o-teste-passar
