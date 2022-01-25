package concorrencia

import (
	"reflect"
	"testing"
	"time"
)

func MockVerificadorWebsite(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func slowStubVerificadorSite(_ string) bool {
	time.Sleep(20 * time.Microsecond)
	return true
}

func TestVerificaWebsites(t *testing.T) {
	t.Run("Teste verifica sites", func(t *testing.T) {
		sites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}
		expected := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		got := VerificaWebsites(MockVerificadorWebsite, sites)

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}

func BenchmarkVerificadorSites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "uma url"
	}
	for i := 0; i < b.N; i++ {
		VerificaWebsites(slowStubVerificadorSite, urls)
	}
}
