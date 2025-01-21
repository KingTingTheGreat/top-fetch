package output

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kingtingthegreat/top-fetch/config"
	"github.com/kingtingthegreat/top-fetch/env"
)

func WriteToFile(outputString string) (string, error) {
	cfg := config.Config()
	outputFile := cfg.File
	if cfg.Path == "relative" {
		basePath, err := env.GetBasePath()
		if err != nil {
			return "", err
		}
		outputFile = filepath.Join(basePath, cfg.File)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return "", fmt.Errorf("could not create output file")
	}
	defer file.Close()

	_, err = file.WriteString(outputString)
	if err != nil {
		return "", fmt.Errorf("could not write to output file")
	}

	return outputFile, nil
}
