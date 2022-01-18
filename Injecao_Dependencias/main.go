package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func SayHello(escritor io.Writer, name string) {
	fmt.Fprintf(escritor, "Olá %s\n", name)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	SayHello(w, "Golang")
}

func main() {
	fmt.Println("Startin server on :3000")
	err := http.ListenAndServe(":3000", http.HandlerFunc(HandlerMeuCumprimento))

	if err != nil {
		log.Printf("error to create server %s\n", err)
	}
}
