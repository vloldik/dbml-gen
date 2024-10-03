package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"go/token"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase" // Added for case conversion
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
)

// GenerateModels generates Go model files from DB tables.
func GenerateModels(parsed *models.DBML, outputDir string, useGorm bool) error {
	for _, table := range parsed.Tables {
		file := generateModelContent(table, useGorm)
		fileName := fmt.Sprintf("%s.go", strings.ToLower(table.Name))
		filePath := filepath.Join(outputDir, fileName)

		err := file.Save(filePath)
		if err != nil {
			return fmt.Errorf("error writing file %s: %v", filePath, err)
		}
	}
	return nil
}

// generateModelContent generates the content of a Go model file for a single table.
func generateModelContent(table *models.Table, useGorm bool) *jen.File {
	file := jen.NewFile("models")

	structFields := make([]jen.Code, 0, len(table.Entries.Columns))

	for _, column := range table.Entries.Columns {
		goFieldName := toExportedGoName(column.Name)
		goType := mapDBTypeToGoType(column.Type)

		field := jen.Id(goFieldName).Id(goType.GoString())

		structFields = append(structFields, field)
	}

	file.Type().Id(toExportedGoName(table.Name)).Struct(structFields...)

	return file
}

// toExportedGoName converts a snake_case string to CamelCase and ensures it's exported.
func toExportedGoName(name string) string {
	camel := strcase.ToCamel(name)
	// Handle Go reserved words by appending an underscore
	if token.Lookup(camel).IsKeyword() {
		camel += "_"
	}
	return camel
}

// toJSONTag converts the column name to a JSON-friendly tag in snake_case.
func toJSONTag(name string) string {
	return name
}

// mapDBTypeToGoType maps database types to Go types.
func mapDBTypeToGoType(dbType string) jen.Statement {
	switch strings.ToLower(dbType) {
	case "int", "integer":
		return *jen.Int()
	case "bigint":
		return *jen.Int64()
	case "varchar", "text", "char":
		return *jen.String()
	case "boolean":
		return *jen.Bool()
	case "float", "double":
		return *jen.Float64()
	case "date", "datetime", "timestamp":
		return *jen.Qual("time", "Time")
	default:
		return *jen.Any()
	}
}
