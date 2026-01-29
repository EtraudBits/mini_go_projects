package main

import (
	"fmt"
	"os"
)

func main() {
	if err := Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler entrada: %v\n", err)
		os.Exit(1)
	}
}