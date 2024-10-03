package models

type Index struct {
	IsPrimaryKey bool
	IsUnique     bool
	Note         string

	Type   string
	Fields []*Field
	Exprs  []string
}
