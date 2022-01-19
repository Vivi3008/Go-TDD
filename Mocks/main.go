package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "vai!"
const inicio = 3

func main() {
	sl := SleeperPadrao{t: (1 * time.Second)}
	Contagem(os.Stdout, &sl)
}

type Sleeper interface {
	Sleep()
}

type SleeperPadrao struct {
	t time.Duration
}

func (s *SleeperPadrao) Sleep() {
	time.Sleep(s.t)
}

//mock
type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicio; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}

//https://larien.gitbook.io/aprenda-go-com-testes/primeiros-passos-com-go/mocks#ainda-temos-alguns-problemas
