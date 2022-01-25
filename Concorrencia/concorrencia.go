package concorrencia

import "time"

type VerificadorWebsites func(string) bool

func VerificaWebsites(vw VerificadorWebsites, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = vw(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}
