package concorrencia

type VerificadorWebsites func(string) bool

type resultado struct {
	string
	bool
}

func VerificaWebsites(vw VerificadorWebsites, urls []string) map[string]bool {
	results := make(map[string]bool)
	canalResultado := make(chan resultado)

	for _, url := range urls {
		go func(u string) {
			canalResultado <- resultado{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		resultado := <-canalResultado
		results[resultado.string] = resultado.bool
	}

	return results
}
