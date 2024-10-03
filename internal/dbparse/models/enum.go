package models

// Enums are currently not supported
type EnumRef struct {
	Name string `"enum:" @Ident`
}

// Enums are currently not supported
type Enum struct {
	Type   string   `"Enum" @Ident "{"`
	Name   string   `@Ident`
	Values []string `@Ident* "}"`
}
