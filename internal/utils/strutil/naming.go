package strutil

import (
	"go/token"
	"strings"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// toExportedGoName converts a snake_case string to CamelCase and ensures it's exported.
func ToExportedGoName(name string) string {
	camel := strcase.ToCamel(name)
	// Handle Go reserved words by appending an underscore
	if token.Lookup(camel).IsKeyword() {
		camel += "_"
	}
	if cutted, ok := strings.CutSuffix(camel, "Id"); ok {
		camel = cutted + "ID"
	}
	return camel
}

// toJSONTag converts the column name to a JSON-friendly tag in snake_case.
func ToJSONTag(name string) string {
	return name
}

func ToSingle(name string) string {
	cli := pluralize.NewClient()
	if cli.IsPlural(name) {
		return cli.Singular(name)
	}
	return name
}

func ToPlural(name string) string {
	cli := pluralize.NewClient()
	if cli.IsSingular(name) {
		return cli.Plural(name)
	}
	return name
}

func CreateManyToManyName(tableNameA, tableNameB string) string {
	tableNameA = ToSingle(tableNameA)
	tableNameB = ToPlural(tableNameB)
	return strcase.ToSnake(tableNameA) + "_" + strcase.ToSnake(tableNameB)
}
