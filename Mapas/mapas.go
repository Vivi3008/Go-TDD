package mapas

import "errors"

type Dicionario map[string]string

var ErrPalavraInexistente = errors.New("palavra inexistente")

func (d Dicionario) Busca(valor string) (string, error) {
	palavra, existe := d[valor]
	if !existe {
		return "", ErrPalavraInexistente
	}
	return palavra, nil
}
