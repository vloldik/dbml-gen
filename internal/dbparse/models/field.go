package models

type Field struct {
	TableName string // For easier relation creation
	Name      string
	Type      string
	Len       int

	IsPrimaryKey bool
	IsIncrement  bool
	IsUnique     bool
	IsNotNull    bool
	Note         string
	DefaultValue string

	Relations []*Relation
	Indexes   []*Index `json:"-"`
}

type Relation struct {
	RelationType RelationType
	SecondTable  string
	SecondField  string
}
