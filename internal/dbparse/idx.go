package dbparse

import (
	"regexp"
	"strings"

	"guthub.com/vloldik/dbml-gen/internal/models"
)

func (p *Parser) parseIndexes(tableContent string) []models.Index {
	var indexes []models.Index

	// Регулярное выражение для извлечения определений индексов
	indexRegex := regexp.MustCompile(`indexes\s*{(.*?)}`)
	indexMatch := indexRegex.FindStringSubmatch(tableContent)

	if len(indexMatch) > 1 {
		indexContent := indexMatch[1]
		indexDefRegex := regexp.MustCompile(`\((.*?)\)`)
		indexDefs := indexDefRegex.FindAllStringSubmatch(indexContent, -1)

		for _, indexDef := range indexDefs {
			columns := strings.Split(indexDef[1], ",")
			for i := range columns {
				columns[i] = strings.TrimSpace(columns[i])
			}
			indexes = append(indexes, models.Index{Columns: columns})
		}
	}

	return indexes
}
