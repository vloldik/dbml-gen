package models

import "encoding/json"

type Field struct {
	Name string
	Type string
	Len  int

	IsPrimaryKey bool
	IsIncrement  bool
	IsUnique     bool
	IsNotNull    bool
	Note         string
	DefaultValue string

	ReferencesTo []*Reference
	ReferencedBy []*Reference
	Indexes      []*Index `json:"-"`
}

type Reference struct {
	RelationType RelationType
	SecondTable  *Table
	SecondField  *Field
}

func (r *Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		RelationType RelationType
		SecondTable  string
		SecondField  string
	}{
		RelationType: r.RelationType,
		SecondTable:  r.SecondTable.Name,
		SecondField:  r.SecondField.Name,
	})
}
