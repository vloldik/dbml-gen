package dbparse

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/converts"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
)

type Parser struct{}

func (p *Parser) Parse(dbml string) (*models.DBML, error) {
	// Replace string
	parserLexer := lexer.MustSimple([]lexer.SimpleRule{
		{Name: "String", Pattern: `"(\\"|\\[\n\r]+|[^"\n\r])*"|'(\\'|\\[\n\r]+|[^"\n\r])*'`},
		{Name: "Comment", Pattern: `(\/\/[^\n\r]*)`},

		{Name: "EOL", Pattern: `[\n\r]+`},

		{Name: "whitespace", Pattern: `[ \t]`},

		{Name: `Ident`, Pattern: `[a-zA-Z_][a-zA-Z0-9_]*`},
		{Name: "Color", Pattern: `\#[\dA-Fa-f]+`},

		{Name: "Punct", Pattern: `<>|[-[!@#$%^&*()+_={}\|:;"'<,>.?\/]|]`},

		{Name: "DBStatement", Pattern: "\\`(\\\\`|[^\\`])*\\`"},
		{Name: "Number", Pattern: `[-+]?([\dA-Fa-f]*\.)?[\dA-Fa-f]+`},
	})

	parser := participle.MustBuild[parseobj.DBML](
		participle.Lexer(parserLexer),
		participle.Elide("whitespace", "EOL", "Comment"),
		participle.CaseInsensitive("Ident"),
		participle.Union[parseobj.Setting](
			&parseobj.HeaderColorSetting{},
			&parseobj.SettingDefaultValue{},
			&parseobj.SettingIncrement{},
			&parseobj.SettingNotNull{},
			&parseobj.SettingNote{},
			&parseobj.SettingPrimaryKey{},
			&parseobj.SettingReference{},
			&parseobj.SettingUnique{},
			&parseobj.SettingsIndexType{},
		),
	)

	parsed, err := parser.ParseString("", dbml)

	if err != nil {
		return nil, err
	}

	converter := converts.NewParseObjectToModelConverter()

	return converter.ObjToModel(parsed)
}

func New() *Parser {
	return &Parser{}
}
