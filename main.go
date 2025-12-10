package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Definir el flag para la profundidad
	// default: 1, descripción: "Nivel de profundidad (max 3)"
	depthPtr := flag.Int("depth", 1, "Nivel de profundidad de la estructura (1-3)")

	// Parsear los flags. Esto debe hacerse antes de leer los argumentos posicionales.
	flag.Parse()

	// Obtener los argumentos que no son flags (la ruta)
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Uso: treeprinter [-depth N] <ruta>")
		fmt.Println("Ejemplo: treeprinter -depth 2 .")
		return
	}

	rootPath := args[0]
	maxDepth := *depthPtr

	// Validación: Si el usuario pone algo mayor a 3, lo limitamos a 3.
	// Si pone algo menor a 1, lo forzamos a 1.
	if maxDepth > 3 {
		fmt.Println("Advertencia: La profundidad máxima permitida es 3. Se ajustará a 3.")
		maxDepth = 3
	} else if maxDepth < 1 {
		maxDepth = 1
	}

	exclude := []string{".git", "node_modules", "dist", ".idea"}

	// Pasamos maxDepth a la función
	err := PrintTree(rootPath, exclude, maxDepth)
	if err != nil {
		log.Fatalf("Error al imprimir estructura: %v", err)
	}
}
