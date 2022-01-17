package mapas

import "errors"

type Dicionario map[string]string

func (d Dicionario) Busca(valor string) (string, error) {
	palavra, existe := d[valor]
	if !existe {
		return "", errors.New("palavra inexistente")
	}
	return palavra, nil
}
