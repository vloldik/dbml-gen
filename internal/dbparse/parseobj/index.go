package parseobj

type IndexSetting struct {
	PrimaryKey bool    `  @("pk" | "primary key")`
	Unique     bool    `| @"unique"?`
	Type       *string `| "type" ":" @Ident`
	Note       *string `| "note" ":" @String`
}
