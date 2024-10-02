package dbparse

import (
	"regexp"
	"strings"

	"guthub.com/vloldik/dbml-gen/internal/models"
)

func (p *Parser) parseColumns(tableContent string) []models.Column {
	var columns []models.Column

	// Регулярное выражение для извлечения определений колонок
	columnRegex := regexp.MustCompile(`(\w+)\s+(\w+)(?:\s*\[(.*?)\])?`)
	columnDefs := columnRegex.FindAllStringSubmatch(tableContent, -1)

	for _, columnDef := range columnDefs {
		column := models.Column{
			Name: columnDef[1],
			Type: columnDef[2],
		}

		if len(columnDef) > 3 && columnDef[3] != "" {
			p.parseFlags(columnDef[3], &column)
		}

		columns = append(columns, column)
	}

	return columns
}

func (p *Parser) parseFlags(flagsString string, column *models.Column) {
	flags := strings.Split(flagsString, ",")
	for _, flag := range flags {
		flag = strings.TrimSpace(flag)
		switch flag {
		case "pk", "primary key":
			column.IsPK = true
		case "unique":
			column.IsUnique = true
		}
		if strings.HasPrefix(flag, "note:") {
			column.Note = strings.Trim(strings.Trim(strings.TrimPrefix(flag, "note:"), "'"), "\"")
		}
	}
}
