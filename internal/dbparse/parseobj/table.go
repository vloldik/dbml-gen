package parseobj

import "github.com/alecthomas/participle/v2/lexer"

type NamespacedName struct {
	Namespace *string `(@Ident ".")?`
	Name      string  `@Ident`
}

type TableContent struct {
	Columns []*Column `@@*`
	Indexes []*Index  `( "indexes" "{" @@* "}" )?`
}

type Column struct {
	Name     string    `@Ident`
	Type     string    `@Ident` // TODO: make []string in future for enum support
	Len      []int     `( "(" @Number ("," @Number )* ")" )?`
	Settings *Settings `@@`
}

type Index struct {
	Tokens   []lexer.Token
	Fields   []string  `(@Ident | ( "(" (@Ident | @DBStatement) ( "," (@Ident | @DBStatement) )* ")" ))`
	Settings *Settings `@@`
}

type ReferenceField struct {
	NameParts []string `@Ident ("." @Ident)+`
}

type EnumValue struct {
	Name     string    `@Ident`
	Settings *Settings `@@`
}
