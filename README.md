# 🗂️ TreePrinter - Visualiza la estructura de archivos en consola

`TreePrinter` es una utilidad escrita en Go que imprime la estructura jerárquica de carpetas y archivos de un proyecto, con soporte para excluir directorios como `node_modules`, `.git`, etc.

## 📦 Estructura del proyecto

```
treeprinter/
├── go.mod
├── main.go
└── treeprinter.go
```

## 🚀 Uso rápido

```bash
go run . <ruta>
```
## Ejemplo

```bash
go run . .
```

## Salida

```
treeprinter/
├── go.mod
├── main.go
└── treeprinter.go
```

## 🛠️ Compilar el ejecutable y Registrar en Bin

```bash
go build -o treeprinter

# Mover el binario a esa carpeta
mkdir -p ~/bin
mv treeprinter ~/bin/

# Agregar ~/bin a tu PATH en ```~/.bashrc``` ó ```~/.zshrc```

export PATH="$HOME/bin:$PATH"

source ~/.bashrc
# o
source ~/.zshrc
```

