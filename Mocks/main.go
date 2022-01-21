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
	sl := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sl)
}

type Sleeper interface {
	Sleep()
}

func (s *SleeperConfiguravel) Sleep() {
	s.pausa(s.duracao)
}

//mock
type SleeperSpy struct {
	Chamadas int
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

const escrita = "escrita"
const pausa = "pausa"

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicio; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}
