package models

import (
	"fmt"
	"strings"

	"github.com/vloldik/dbml-gen/internal/utils/hashutil"
	"github.com/vloldik/dbml-gen/internal/utils/listutil"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

type Table struct {
	Fields  []*Field
	Indexes []*Index

	TableName *NamespacedName
	Alias     *string
	Note      string
}

func (t Table) GetFieldByName(name string) (*Field, error) {
	val := listutil.SearchFunc(t.Fields, func(f *Field, i int) bool { return f.DBName == name })

	if val == nil {
		return nil, fmt.Errorf("field %s.%s not found", t.TableName, name)
	}

	return val, nil
}

func (t Table) GetFieldByDisplayName(name string) (*Field, error) {
	val := listutil.SearchFunc(t.Fields, func(f *Field, i int) bool { return f.DisplayName() == name })
	if val == nil {
		return nil, fmt.Errorf("field %s.%s not found", t.DisplayName(), name)
	}

	return val, nil
}

func (t Table) DisplayName() string {
	return strutil.ToSingle(strutil.ToExportedGoName(t.TableName.BaseName))
}

func (t Table) PackageName() string {
	return strings.ReplaceAll(t.TableName.Namespace, "-", "_")
}

func (t Table) Hash() uint32 {
	return hashutil.FnvSumm(
		[]byte(t.TableName.FullName()),
	)
}

func (t Table) NotePrepared() string {
	return strutil.TryUnquote(t.Note)
}
