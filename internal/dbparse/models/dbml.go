package models

type DBML struct {
	Tables    []*Table
	Relations []*Relationship `json:"-"`
	// Enums
}
