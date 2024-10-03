package indexer

import (
	"fmt"

	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

type tableMap map[string]*models.Table

type Indexer struct {
	tableMap tableMap
}

func NewIndexCreator(tableMap map[string]*models.Table) *Indexer {
	return &Indexer{
		tableMap: tableMap,
	}
}

func (i *Indexer) CreateIndexes(table *parseobj.Table) ([]*models.Index, error) {
	if table.Content == nil {
		return []*models.Index{}, nil
	}

	indexList := make([]*models.Index, 0)
	for _, index := range table.Content.Indexes {
		indexModel, err := i.indexFromFields(table.Name, index.Fields)
		if err != nil {
			return nil, err
		}

		indexList = append(indexList, i.applyIndexSettings(indexModel, index.Settings))
	}

	return indexList, nil
}

func (i *Indexer) indexFromFields(tableName string, fields []string) (*models.Index, error) {
	idx := &models.Index{}

	table, ok := i.tableMap[tableName]

	if !ok {
		return nil, fmt.Errorf("table %s not found", tableName)
	}

	var fieldString string

	for _, field := range fields {
		if field[0] == '`' {
			idx.Exprs = append(idx.Exprs, field)
			continue
		} else if field[0] == '"' {
			inner, err := strutil.UnquoteString(field)
			if err != nil {
				return nil, err
			}

			fieldString = inner
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

func (i *Indexer) applyIndexSettings(index *models.Index, settings []*parseobj.IndexSetting) *models.Index {
	for _, setting := range settings {
		if setting.Note != nil {
			index.Note = *setting.Note
		} else if setting.PrimaryKey {
			index.IsPrimaryKey = true
		} else if setting.Unique {
			index.IsUnique = true
		} else if setting.Type != nil {
			index.Type = *setting.Type
		}
	}

	return index
}
