package models

type IndexSetting struct {
	PrimaryKey *string `  @("pk" | "primary key")`
	Unique     *bool   `| @"unique"?`
	Type       *string `| "type" ":" @Ident`
	Note       *string `| "note" ":" @String`
}
