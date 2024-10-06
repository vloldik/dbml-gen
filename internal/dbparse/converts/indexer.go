package converts

import (
	"fmt"

	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

func (c *ParseObjectToModelConverter) CreateIndexes(table *parseobj.StructureTable) ([]*models.Index, error) {
	if table.Content == nil {
		return []*models.Index{}, nil
	}

	indexList := make([]*models.Index, 0)
	for _, index := range table.Content.Indexes {
		indexModel, err := c.indexFromFields(
			models.NewNamespacedName(table.Name.Namespace, table.Name.Name).FullName(), index.Fields,
		)
		if err != nil {
			return nil, err
		}
		if err := c.applySettings(indexModel, index.Settings); err != nil {
			return nil, err
		}

		indexList = append(indexList, indexModel)
	}

	return indexList, nil
}

func (c *ParseObjectToModelConverter) indexFromFields(tableName string, fields []string) (*models.Index, error) {
	idx := &models.Index{}

	table, ok := c.tableMap[tableName]

	if !ok {
		return nil, fmt.Errorf("table %s not found", tableName)
	}

	var fieldString string

	for _, field := range fields {
		if field[0] == '`' {
			idx.Exprs = append(idx.Exprs, field)
			continue
		} else if field[0] == '"' || field[0] == '\'' {
			inner, err := strutil.UnquoteString(field)
			if err != nil {
				return nil, err
			}

			fieldString = inner
		} else {
			fieldString = field
		}

		fieldModel, err := table.GetFieldByName(fieldString)

		if err != nil {
			return nil, err
		}

		idx.Fields = append(idx.Fields, fieldModel)
		fieldModel.Indexes = append(fieldModel.Indexes, idx)
	}

	return idx, nil
}
