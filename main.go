package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . <ruta>")
		return
	}

	rootPath := os.Args[1]
	exclude := []string{".git", "node_modules", "dist", ".idea"}

	err := PrintTree(rootPath, exclude)
	if err != nil {
		log.Fatalf("Error al imprimir estructura: %v", err)
	}
}
