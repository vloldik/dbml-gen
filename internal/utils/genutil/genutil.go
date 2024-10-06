package genutil

import (
	"fmt"
	"strings"

	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

func CreateBasicGORMTags(field *models.Field) ([]string, error) {
	// Добавляем теги GORM
	tags := make([]string, 1)
	tags[0] = fmt.Sprintf("column:%s", field.Name)

	if field.IsPrimaryKey {
		tags = append(tags, "primaryKey")
	}
	if field.IsUnique {
		tags = append(tags, "unique")
	}
	if field.Type != "" {
		gormType, needToSpecify := getGORMTypeForName(field.Type)
		if needToSpecify {
			tags = append(tags, fmt.Sprintf("type:%s", gormType))
		}
	}
	if field.IsNotNull {
		tags = append(tags, "not null")
	}
	if field.IsIncrement {
		tags = append(tags, "autoIncrement")
	}
	if field.Len != 0 {
		tags = append(tags, fmt.Sprintf("size:%d", field.Len))
	}
	if field.DefaultValue != "" {
		defaultVal, removed := strutil.RemoveQuotes(field.DefaultValue, "\"")
		if !removed {
			defaultVal, _ = strutil.RemoveQuotes(defaultVal, "'")
		}
		defaultVal = strings.ReplaceAll(defaultVal, "`", "")

		tags = append(tags, fmt.Sprintf("default:%s", defaultVal))
	}

	return tags, nil
}
