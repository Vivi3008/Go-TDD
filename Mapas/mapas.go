package mapas

type Dicionario map[string]string

type ErrDicionario string

const (
	ErrPalavraInexistente ErrDicionario = "palavra inexistente"
	ErrPalavraJaExiste    ErrDicionario = "essa palavra já existe"
	ErrAtualizaPalavra    ErrDicionario = "nao é possivel atualizar, palavra inexistente"
)

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(valor string) (string, error) {
	palavra, existe := d[valor]
	if !existe {
		return "", ErrPalavraInexistente
	}
	return palavra, nil
}

func (d Dicionario) Atualiza(chave, valor string) error {
	_, err := d.Busca(chave)

	if err == ErrPalavraInexistente {
		return ErrAtualizaPalavra
	}

	d[chave] = valor
	return nil
}

func (d Dicionario) Adiciona(chave, valor string) error {
	_, err := d.Busca(chave)

	if err == nil {
		return ErrPalavraJaExiste
	}

	d[chave] = valor
	return nil
}

func (d Dicionario) Deleta(chave string) {
	delete(d, chave)
}
