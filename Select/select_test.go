package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("Teste corredor", func(t *testing.T) {
		servidorLento := CriarServidorComAtraso(20 * time.Millisecond)
		servidorRapido := CriarServidorComAtraso(0 * time.Millisecond)

		defer servidorLento.Close()
		defer servidorRapido.Close()

		urlRapida := servidorRapido.URL
		urlDevagar := servidorLento.URL

		expected := urlRapida
		got := Corredor(urlDevagar, urlRapida)

		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}

func CriarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		rw.WriteHeader(http.StatusOK)
	}))
}
