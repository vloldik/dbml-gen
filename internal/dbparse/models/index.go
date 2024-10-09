package models

import (
	"strings"

	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

type Index struct {
	IsPrimaryKey bool
	IsUnique     bool
	Name         string
	Note         string

	// used to cache calculated name
	cachedName string
	Type       string
	Fields     []*Field
	Exprs      []string
}

func (i *Index) getFirstPart() string {
	if i.IsPrimaryKey {
		return "pk"
	}
	if i.IsUnique {
		return "ux"
	}
	return "ix"
}

func (i *Index) GetName() string {
	if i.cachedName != "" {
		return i.cachedName
	}
	if i.Name != "" {
		return i.Name
	}
	nameBuilder := &strings.Builder{}
	nameBuilder.WriteString(i.getFirstPart())
	nameBuilder.WriteString("_")
	for i, field := range i.Fields {
		if i == 0 {
			nameBuilder.WriteString(strutil.ToSingle(field.Table.TableName.BaseName))
		}
		nameBuilder.WriteString("_")
		nameBuilder.WriteString(field.DBName)
		nameBuilder.WriteString("_")
	}
	i.cachedName = strings.TrimSuffix(nameBuilder.String(), "_")
	return i.cachedName
}
