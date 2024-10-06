package strutil

import (
	"strings"
)

func ConcantatePaths(paths ...string) string {
	if len(paths) == 0 {
		return ""
	}
	if len(paths) == 1 {
		return paths[0]
	}

	result := paths[0]
	for _, path := range paths[1:] {
		path = NormalizePath(path)
		result += ("/" + path)
	}

	return result
}

func NormalizePath(path string) string {
	path = strings.ReplaceAll(path, "\\", "/")
	path = strings.Trim(path, ".")
	path = strings.Trim(path, "/")
	return path
}
