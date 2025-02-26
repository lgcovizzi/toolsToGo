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

// ReadFile lê o conteúdo de um arquivo. O nome do arquivo é obrigatório e o diretório é opcional.
func ReadFile(fileName string, dirPath ...string) (string, error) {
	var filePath string

	if len(dirPath) > 0 {
		filePath = filepath.Join(dirPath[0], fileName)
	} else {
		var err error
		dirPath[0], err = GetProjectRoot()
		if err != nil {
			return "", err
		}
		filePath = filepath.Join(dirPath[0], fileName)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadAllFilesWithExtension retorna uma lista de arquivos com a extensão especificada. O diretório é opcional.
func ReadAllFilesWithExtension(extensao string, dirPath ...string) ([]string, error) {
	var path string

	if len(dirPath) > 0 {
		path = dirPath[0]
	} else {
		var err error
		path, err = GetProjectRoot()
		if err != nil {
			return nil, err
		}
	}

	var arquivosComExtensao []string
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(filePath) == extensao {
			arquivosComExtensao = append(arquivosComExtensao, filePath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return arquivosComExtensao, nil
}
