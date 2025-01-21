package env

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetBasePath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("cannot get executable path")
	}

	execDir := filepath.Dir(execPath)
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("cannot get current working directory")
	}

	if !strings.HasPrefix(execDir, os.TempDir()) {
		return execDir, nil
	}
	return cwd, nil
}
