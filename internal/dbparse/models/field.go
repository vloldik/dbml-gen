package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vloldik/dbml-gen/internal/utils/genutil"
	"github.com/vloldik/dbml-gen/internal/utils/hashutil"
	"github.com/vloldik/dbml-gen/internal/utils/listutil"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

type Field struct {
	Table  *Table // backref
	DBName string

	Type string
	Len  []int

	IsPrimaryKey bool
	IsIncrement  bool
	IsUnique     bool
	IsNotNull    bool
	Note         string
	DefaultValue string

	Indexes []*Index `json:"-"`
}

type FieldRelation struct {
	RelationType RelationType
	SecondTable  *NamespacedName
	SecondField  string
}

func (f *Field) DisplayName() string {
	return strutil.ToExportedGoName(f.DBName)
}

func (f Field) Hash() uint32 {
	return hashutil.FnvSumm(
		[]byte(f.DBName),
	) + f.Table.Hash()
}

func (f *Field) CreateBasicGORMTags() ([]string, error) {
	// Добавляем теги GORM
	tags := make([]string, 1)
	tags[0] = fmt.Sprintf("column:%s", f.DBName)

	if f.IsPrimaryKey {
		tags = append(tags, "primaryKey")
	}
	if f.IsUnique {
		tags = append(tags, "unique")
	}
	if f.Type != "" {
		gormType, needToSpecify := genutil.GetGORMTypeForName(f.Type)
		if len(f.Len) == 1 {
			tags = append(tags, fmt.Sprintf("size:%d", f.Len[0]))
		} else if len(f.Len) > 1 {
			gormType += fmt.Sprintf("(%s)", strings.Join(listutil.Map(f.Len, func(len int, _ int) string {
				return strconv.Itoa(len)
			}), ","))
		}
		if needToSpecify {
			tags = append(tags, fmt.Sprintf("type:%s", gormType))
		}
	}
	if f.IsNotNull {
		tags = append(tags, "not null")
	}
	if f.IsIncrement {
		tags = append(tags, "autoIncrement")
	}
	if f.DefaultValue != "" {
		defaultVal, removed := strutil.RemoveQuotes(f.DefaultValue, "\"")
		if !removed {
			defaultVal, _ = strutil.RemoveQuotes(defaultVal, "'")
		}
		defaultVal = strings.ReplaceAll(defaultVal, "`", "")

		tags = append(tags, fmt.Sprintf("default:%s", defaultVal))
	}

	return tags, nil
}

func (f Field) NotePrepared() string {
	return strutil.TryUnquote(f.Note)
}
