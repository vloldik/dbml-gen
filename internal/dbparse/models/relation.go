package models

import (
	"hash/fnv"
)

type RelationType int8

const (
	OneToOne RelationType = iota
	ManyToOne
	ManyToMany
)

type Relationship struct {
	RelationType RelationType

	FromTable *Table
	FromField *Field

	ToField *Field
	ToTable *Table
}

// Calculating hash function for easy comparation
func (r Relationship) Hash() uint32 {
	h := fnv.New32a()

	// Type does not matter, there can be only one relation type
	// beetween two fields

	h.Write([]byte(r.FromTable.Name))
	h.Write([]byte(r.FromField.Name))

	h.Write([]byte(r.ToTable.Name))
	h.Write([]byte(r.ToField.Name))

	return h.Sum32()
}
