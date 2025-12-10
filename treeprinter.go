package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// PrintTree imprime la estructura de archivos desde rootPath hasta una profundidad máxima (maxDepth).
// maxDepth define cuántos niveles de subdirectorios se mostrarán.
func PrintTree(rootPath string, excludeDirs []string, maxDepth int) error {
	excluded := make(map[string]bool)
	for _, dir := range excludeDirs {
		excluded[dir] = true
	}

	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return fmt.Errorf("error resolviendo ruta absoluta: %w", err)
	}

	rootName := filepath.Base(absPath)
	fmt.Println(rootName + "/")

	// Iniciamos la recursión con currentDepth en 0
	return printRecursive(absPath, "", excluded, 0, maxDepth)
}

func printRecursive(path string, prefix string, excluded map[string]bool, currentDepth int, maxDepth int) error {
	// Si hemos alcanzado la profundidad máxima, dejamos de explorar subcarpetas
	if currentDepth >= maxDepth {
		return nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		isLast := i == len(entries)-1
		connector := "├── "
		subPrefix := "│   "
		if isLast {
			connector = "└── "
			subPrefix = "    "
		}

		fmt.Println(prefix + connector + entry.Name())

		if entry.IsDir() {
			if excluded[entry.Name()] {
				continue
			}
			newPath := filepath.Join(path, entry.Name())
			// Aumentamos currentDepth en 1 para la siguiente llamada recursiva
			err := printRecursive(newPath, prefix+subPrefix, excluded, currentDepth+1, maxDepth)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
