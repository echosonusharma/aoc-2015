package utils

import (
	"os"
	"path/filepath"
)

func LoadFile(fileName string) (string, error) {
	baseDir := "inputs"
	filePath := filepath.Join(baseDir, fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
