package genutil

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

// mapDBTypeToGoType maps database types to Go types.
func MapDBTypeToGoType(statement *jen.Statement, dbType string) {
	switch strings.ToLower(dbType) {
	case "int", "integer":
		statement.Int()
	case "bigint":
		statement.Int64()
	case "varchar", "text", "char":
		statement.String()
	case "boolean":
		statement.Bool()
	case "float", "double":
		statement.Float64()
	case "date", "datetime", "timestamp":
		statement.Qual("time", "Time")
	default:
		statement.Any()
	}
}

func CreateGORMTags(field *models.Field) (map[string]string, error) {
	// Добавляем теги GORM
	tags := make([]string, 0)
	tags = append(tags, fmt.Sprintf("column:%s", field.Name))

	if field.IsPrimaryKey {
		tags = append(tags, "primaryKey")
	}
	if field.IsUnique {
		tags = append(tags, "unique")
	}
	if field.IsNotNull {
		tags = append(tags, "not null")
	}
	if field.DefaultValue != "" {
		defaultVal, removed := strutil.RemoveQuotes(field.DefaultValue, "\"")
		if !removed {
			defaultVal, _ = strutil.RemoveQuotes(defaultVal, "'")
		}
		defaultVal = strings.ReplaceAll(defaultVal, "`", "")

		tags = append(tags, fmt.Sprintf("default:%s", defaultVal))
	}

	if len(tags) == 0 {
		return map[string]string{}, nil
	}

	return map[string]string{"gorm": strings.Join(tags, ";")}, nil
}
