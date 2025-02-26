package toolsToGo

import (
	"fmt"
	"os"
	"path/filepath"
)

// SayHello prints a greeting
func SayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// GetProjectRoot returns the root directory of the project
func GetProjectRoot() (string, error) {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Navigate to the project root
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
	return "", fmt.Errorf("project root not found")
}

// ReadFile reads the content of a file. The file name is required and the directory is optional.
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

// ReadAllFilesWithExtension returns a list of files with the specified extension. The directory is optional.
func ReadAllFilesWithExtension(extension string, dirPath ...string) ([]string, error) {
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

	var filesWithExtension []string
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(filePath) == extension {
			filesWithExtension = append(filesWithExtension, filePath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return filesWithExtension, nil
}
