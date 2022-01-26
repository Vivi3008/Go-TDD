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
		got, _ := Corredor(urlDevagar, urlRapida)

		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})

	t.Run("Deve retornar um erro se a requisicao demorar mais de 10s", func(t *testing.T) {
		servidorA := CriarServidorComAtraso(25 * time.Millisecond)

		defer servidorA.Close()

		_, err := Configuravel(servidorA.URL, servidorA.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("Expected error if request duration is greater than 10 seconds")
		}

	})
}

func CriarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(atraso)
		rw.WriteHeader(http.StatusOK)
	}))
}
