package strutil

import (
	"go/token"

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
