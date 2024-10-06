package strutil

import (
	"go/token"
	"strings"

	"github.com/iancoleman/strcase"
)

// toExportedGoName converts a snake_case string to CamelCase and ensures it's exported.
func ToExportedGoName(name string) string {
	camel := strcase.ToCamel(name)
	// Handle Go reserved words by appending an underscore
	if token.Lookup(camel).IsKeyword() {
		camel += "_"
	}
	return camel
}

// toJSONTag converts the column name to a JSON-friendly tag in snake_case.
func ToJSONTag(name string) string {
	return name
}

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
