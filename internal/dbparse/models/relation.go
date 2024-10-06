package models

import (
	"fmt"
	"hash/fnv"
	"strconv"

	"github.com/vloldik/dbml-gen/internal/dbparse/parseobj"
)

type RelationType int8

const (
	OneToOne RelationType = iota
	OneToMany
	ManyToOne
	ManyToMany
)

func (r RelationType) Name() string {
	name, ok := map[RelationType]string{
		OneToOne:   "One to one",
		OneToMany:  "One to many",
		ManyToOne:  "Many to one",
		ManyToMany: "Many to many",
	}[r]
	if !ok {
		return "unknown"
	}
	return name
}

func (r *RelationType) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(r.Name())), nil
}

func RelationTypeFromParsed(parsed *parseobj.RelationshipType) (RelationType, error) {
	if parsed.ManyToMany {
		return ManyToMany, nil
	} else if parsed.ManyToOne {
		return ManyToOne, nil
	} else if parsed.OneToMany {
		return OneToMany, nil
	} else if parsed.OneToOne {
		return OneToOne, nil
	}

	return -1, fmt.Errorf("unknown relations type")
}

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

	h.Write([]byte(r.FromTable.Name.FullName()))
	h.Write([]byte(r.FromField.Name))

	h.Write([]byte(r.ToTable.Name.FullName()))
	h.Write([]byte(r.ToField.Name))

	return h.Sum32()
}
