package parseobj

type DBML struct {
	Tables []*Table         `(@@)*`
	Enums  []*Enum          `(@@)*`
	Refs   []*FullReference `(@@)*`
}

type Table struct {
	Name     string    `"Table" @Ident`
	Settings *Settings `@@`

	Content *TableContent `"{"  @@? "}"`
}

type TableContent struct {
	Columns []*Column `@@*`
	Indexes []*Index  `( "indexes" "{" @@* "}" )?`
}

type Column struct {
	Name     string    `@Ident`
	Type     string    `@Ident`
	Len      int       `("("@Number")")?`
	Settings *Settings `@@`
}

type Index struct {
	Fields   []string  `@Ident | ( "(" (@Ident | @DBStatement) ( "," (@Ident | @DBStatement) )* ","? ")")`
	Settings *Settings `@@`
}

type FullReference struct {
	Field            *ReferenceField   `"ref" ":" @@`
	Type             *RelationshipType `@@`
	ReferenceToField *ReferenceField   `@@`
}

type ReferenceField struct {
	Table  string `@Ident`
	Column string `"." @Ident`
}
