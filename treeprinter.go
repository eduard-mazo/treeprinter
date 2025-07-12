package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// PrintTree imprime la estructura de archivos desde rootPath, excluyendo carpetas.
func PrintTree(rootPath string, excludeDirs []string) error {
	excluded := make(map[string]bool)
	for _, dir := range excludeDirs {
		excluded[dir] = true
	}

	absPath, err := filepath.Abs(rootPath)
	if err != nil {
		return fmt.Errorf("error resolviendo ruta absoluta: %w", err)
	}

	rootName := filepath.Base(absPath)
	fmt.Println(rootName + "/") // ✅ Ahora imprimirá correctamente el nombre del directorio raíz

	return printRecursive(absPath, "", excluded)
}

func printRecursive(path string, prefix string, excluded map[string]bool) error {
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
			err := printRecursive(newPath, prefix+subPrefix, excluded)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
