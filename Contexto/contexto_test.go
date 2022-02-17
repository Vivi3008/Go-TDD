package contexto

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyResponseWriter struct {
	writenn bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.writenn = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.writenn = true
	return 0, errors.New("nao implementado")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.writenn = true
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("Spy store foi cancelado")
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func TestContexto(t *testing.T) {
	var data = "Hello world"
	t.Run("Contexto", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("Expected %s, got %s", data, response.Body.String())
		}
	})
	t.Run("avisa store para cancelar o trabalho se a request for cancelada", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}
		svr.ServeHTTP(response, request)

		if response.writenn {
			t.Error("uma resposta nao deveria ter sido escrita")
		}
	})
}
