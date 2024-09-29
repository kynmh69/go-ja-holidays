package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("project root not found")
		}
		currentDir = parentDir
	}
}

// JoinProjectRootPath プロジェクトルートにファイルパスを結合する
func JoinProjectRootPath(paths ...string) (string, error) {
	root, err := GetProjectRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(append([]string{root}, paths...)...), nil
}
