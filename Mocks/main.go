package main

import (
	"fmt"
	"io"
	"os"
)

const ultimaPalavra = "vai!"
const inicio = 3

func main() {

	Contagem(os.Stdout)
}

func Contagem(saida io.Writer) {
	for i := inicio; i >= 1; i-- {
		fmt.Fprintln(saida, i)
	}
	fmt.Fprint(saida, ultimaPalavra)
}
