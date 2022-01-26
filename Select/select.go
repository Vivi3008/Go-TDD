package selects

import (
	"fmt"
	"net/http"
	"time"
)

var tempoLimiteDez = 10 * time.Second

func Corredor(url1, url2 string) (string, error) {
	return Configuravel(url1, url2, tempoLimiteDez)
}

func Configuravel(url1, url2 string, tempoLimite time.Duration) (string, error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(tempoLimite):
		return "", fmt.Errorf("request took more than 10 seconds")
	}
}

func ping(url string) chan bool {
	canal := make(chan bool)

	go func(u string) {
		http.Get(u)
		canal <- true
	}(url)

	return canal
}
