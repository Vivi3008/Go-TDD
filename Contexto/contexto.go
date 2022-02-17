package contexto

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := store.Fetch(r.Context())
		if err != nil {
			fmt.Printf("error to request: %v\n", err)
			return
		}
		fmt.Fprint(w, result)
	}
}
