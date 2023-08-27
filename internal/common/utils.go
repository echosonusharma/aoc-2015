package utils

import (
	"io/ioutil"
	"path/filepath"
)

func LoadFile(fileName string) (string, error) {
	baseDir := "inputs"
	filePath := filepath.Join(baseDir, fileName)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
