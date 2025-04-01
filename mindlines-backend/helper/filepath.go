package helper

import (
	"os"
	"path/filepath"
)

func getProjectRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		// Check if "go.mod" exists (adjust this if needed)
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd
		}

		// Move up one directory
		parent := filepath.Dir(wd)
		if parent == wd {
			// Reached the root of the filesystem
			panic("project root not found")
		}
		wd = parent
	}
}

func GetFilePath(relativePath string) string {
	return filepath.Join(getProjectRoot(), relativePath)
}
