package parseobj

type FieldSetting struct {
	PrimaryKey   bool          `  @("pk" | ("primary" "key"))`
	Increment    bool          `| @"increment"`
	Unique       bool          `| @"unique"`
	NotNull      bool          `| @"not_null"`
	Reference    *Relationship `| @@`
	Check        *Check        `| @@`
	EnumRef      *EnumRef      `| @@`
	Note         *string       `| "note" ":" @String`
	DefaultValue *string       `| "default" ":" (@String | @Number | @Ident | @DBStatement)`
}

type Relationship struct {
	Type   *RelationshipType `"ref" ":" @@`
	Table  string            `@Ident`
	Column string            `"." @Ident`
}

type RelationshipType struct {
	OneToOne   bool `@"-"`
	ManyToMany bool `|@"<>"`
	ManyToOne  bool `|@(">" | "<")`
}

type Check struct {
	Expr string `"check(" @String ")"`
}
