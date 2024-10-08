package genutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vloldik/dbml-gen/internal/dbparse/models"
	"github.com/vloldik/dbml-gen/internal/utils/listutil"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
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
		if len(field.Len) == 1 {
			tags = append(tags, fmt.Sprintf("size:%d", field.Len[0]))
		} else if len(field.Len) > 1 {
			gormType += fmt.Sprintf("(%s)", strings.Join(listutil.Map(field.Len, func(len int, _ int) string {
				return strconv.Itoa(len)
			}), ","))
		}
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

func GormTagsFromList(tags ...string) map[string]string {
	return map[string]string{"gorm": strings.Join(tags, ";")}
}
