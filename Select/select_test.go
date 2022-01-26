package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	servidorLento := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Second)
		rw.WriteHeader(http.StatusOK)
	}))

	servidorRapido := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Millisecond)
		rw.WriteHeader(http.StatusOK)
	}))

	t.Run("Teste corredor", func(t *testing.T) {
		urlRapida := servidorRapido.URL
		urlDevagar := servidorLento.URL

		expected := urlRapida
		got := Corredor(urlDevagar, urlRapida)

		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
		servidorLento.Close()
		servidorRapido.Close()
	})
}
