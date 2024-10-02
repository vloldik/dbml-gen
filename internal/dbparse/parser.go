package dbparse

import (
	"regexp"

	"guthub.com/vloldik/dbml-gen/internal/models"
)

type Parser struct{}

func (p *Parser) Parse(dbml string) ([]models.Table, error) {
	var tables []models.Table

	// Регулярное выражение для извлечения определений таблиц
	tableRegex := regexp.MustCompile(`(?s)Table\s+(\w+)\s*{(.*?)}`)
	tableDefs := tableRegex.FindAllStringSubmatch(dbml, -1)

	for _, tableDef := range tableDefs {
		tableName := tableDef[1]
		tableContent := tableDef[2]

		table := models.Table{
			Name:    tableName,
			Columns: p.parseColumns(tableContent),
			Indexes: p.parseIndexes(tableContent),
		}

		tables = append(tables, table)
	}

	return tables, nil
}

func New() *Parser {
	return &Parser{}
}
