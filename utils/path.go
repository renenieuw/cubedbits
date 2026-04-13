package utils

import (
	"path/filepath"
	"strings"
)

// SplitPath takes a path string and returns a slice of all directory names and the filename.
func SplitPath(path string) ([]string, string) {
	// Standardize path separators to forward slashes for consistent splitting
	standardizedPath := filepath.ToSlash(path)
	
	dir, file := filepath.Split(standardizedPath)
	
	// Trim the trailing slash from dir if it exists
	dir = strings.TrimSuffix(dir, "/")
	
	if dir == "" {
		return []string{}, file
	}
	
	// Split the directory part into individual directory names
	parts := strings.Split(dir, "/")
	
	// Filter out any empty strings that might result from multiple slashes or leading slashes
	var directories []string
	for _, part := range parts {
		if part != "" {
			directories = append(directories, part)
		}
	}
	
	return directories, file
}
