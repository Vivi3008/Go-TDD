package mapas

import "errors"

type Dicionario map[string]string

var (
	ErrPalavraInexistente = errors.New("palavra inexistente")
	ErrPalavraJaExiste    = errors.New("essa palavra já existe")
)

func (d Dicionario) Busca(valor string) (string, error) {
	palavra, existe := d[valor]
	if !existe {
		return "", ErrPalavraInexistente
	}
	return palavra, nil
}

func (d Dicionario) Adiciona(chave, valor string) error {
	_, err := d.Busca(chave)

	if err == nil {
		return ErrPalavraJaExiste
	}

	d[chave] = valor
	return nil
}
