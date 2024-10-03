package genutil

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

// mapDBTypeToGoType maps database types to Go types.
func MapDBTypeToGoType(dbType string) *jen.Statement {
	switch strings.ToLower(dbType) {
	case "int", "integer":
		return jen.Int()
	case "bigint":
		return jen.Int64()
	case "varchar", "text", "char":
		return jen.String()
	case "boolean":
		return jen.Bool()
	case "float", "double":
		return jen.Float64()
	case "date", "datetime", "timestamp":
		return jen.Qual("time", "Time")
	default:
		return jen.Any()
	}
}
