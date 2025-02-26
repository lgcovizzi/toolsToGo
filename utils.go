package toolsToGo

import (
	"fmt"
	"os"
	"path/filepath"
)

// SayHello imprime uma saudação
func SayHello(name string) {
	fmt.Printf("Olá, %s!\n", name)
}

func GetProjectRoot() (string, error) {
	// Obtém o diretório atual
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Navega até a raiz do projeto
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return "", fmt.Errorf("raiz do projeto não encontrada")
}
