package models

type Index struct {
	IsPrimaryKey bool
	IsUnique     bool
	Name         string
	Note         string

	Type   string
	Fields []*Field
	Exprs  []string
}
