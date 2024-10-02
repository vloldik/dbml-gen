package models

type Table struct {
	Name    string
	Columns []Column
	Indexes []Index
}

type Column struct {
	Name     string
	Type     string
	IsPK     bool
	IsUnique bool
	Note     string
}

type Index struct {
	Columns []string
}
