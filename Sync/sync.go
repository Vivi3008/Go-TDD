package sync

import "sync"

type Contador struct {
	mu    sync.Mutex
	valor int
}

func (c *Contador) Incrementa() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (v *Contador) Valor() int {
	return v.valor
}
