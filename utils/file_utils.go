package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetAbsoluteFilePath(partialDir string) string {
	// Convert the partial directory to absolute directory
	absDir, err := filepath.Abs(partialDir)
	if err != nil {
		log.Fatalf("Error converting partial directory to absolute: %v", err)
	}

	// Check if the absolute directory exists
	_, err = os.Stat(absDir)
	if os.IsNotExist(err) {
		log.Fatalf("Directory %s does not exist", absDir)
	}
	return absDir
}

func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}
