package selects

import (
	"net/http"
	"time"
)

func Corredor(url1, url2 string) string {
	duracaoUrl1 := medirTempoResposta(url1)
	duracaoUrl2 := medirTempoResposta(url2)

	if duracaoUrl1 > duracaoUrl2 {
		return url2
	}
	return url1
}

func medirTempoResposta(url string) time.Duration {
	inicio := time.Now()
	http.Get(url)
	return time.Since(inicio)
}
