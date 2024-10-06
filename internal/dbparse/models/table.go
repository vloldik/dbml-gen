package models

import "fmt"

type Table struct {
	Fields  []*Field
	Indexes []*Index

	Name  *NamespacedName
	Alias *string
	Note  string
}

func (t Table) GetFieldByName(name string) (*Field, error) {
	for _, field := range t.Fields {
		if field.Name == name {
			return field, nil
		}
	}

	return nil, fmt.Errorf("field %s.%s not found", t.Name, name)
}
