package models

type Field struct {
	TableName *NamespacedName // For easier relation creation
	Name      string
	Type      string
	Len       []int

	IsPrimaryKey bool
	IsIncrement  bool
	IsUnique     bool
	IsNotNull    bool
	Note         string
	DefaultValue string

	Relations []*FieldRelation
	Indexes   []*Index `json:"-"`
}

type FieldRelation struct {
	RelationType RelationType
	SecondTable  *NamespacedName
	SecondField  string
}
