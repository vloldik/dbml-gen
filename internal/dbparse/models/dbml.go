package models

import "github.com/vloldik/dbml-gen/internal/utils/listutil"

type DBML struct {
	Tables    []*Table
	Relations map[uint32][]*Relationship
	// Enums
}

func (dbml DBML) GetTableByName(fullName string) *Table {
	return listutil.SearchFunc(dbml.Tables, func(t *Table, _ int) bool { return t.TableName.FullName() == fullName })
}

func (dbml DBML) RelationsByFieldHash(hash uint32) []*Relationship {
	relations, ok := dbml.Relations[hash]
	if !ok {
		return make([]*Relationship, 0)
	}

	return relations
}
