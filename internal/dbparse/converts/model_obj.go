package converts

import (
	"guthub.com/vloldik/dbml-gen/internal/dbparse/converts/indexer"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/converts/relationer"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
	"guthub.com/vloldik/dbml-gen/internal/utils/maputil"
)

type ParseObjectToModelConverter struct {
	tables          tableMap
	indexCreator    *indexer.Indexer
	relationCreator *relationer.RelationCreator
}

func NewParseObjectToModelConverter() *ParseObjectToModelConverter {
	tables := make(tableMap)
	return &ParseObjectToModelConverter{
		tables:          tables,
		indexCreator:    indexer.NewIndexCreator(tables),
		relationCreator: relationer.NewRelationCreator(tables),
	}
}

type tableMap map[string]*models.Table

// This function converts parsed DBML to model and also checks logic
func (p *ParseObjectToModelConverter) ObjToModel(obj *parseobj.DBML) (*models.DBML, error) {
	p.fillTablesFromDBML(obj.Tables)

	relations, err := p.relationCreator.CreateRelations(obj)
	if err != nil {
		return nil, err
	}

	return &models.DBML{
		Tables:    maputil.Values(p.tables),
		Relations: relations,
	}, nil
}

func (p *ParseObjectToModelConverter) fillTablesFromDBML(from []*parseobj.Table) {
	for _, table := range from {
		p.tables[table.Name] = &models.Table{
			Name:     table.Name,
			Settings: tableSettingsToMap(table.Settings),
			Fields:   createFields(table.Content),
		}
	}
}

func tableSettingsToMap(settings []*parseobj.TableSetting) map[string]string {
	settingMap := make(map[string]string)
	for _, setting := range settings {
		settingMap[setting.Key] = setting.Value
	}
	return settingMap
}

func createFields(content *parseobj.TableContent) []*models.Field {
	fieldList := make([]*models.Field, 0)
	for _, field := range content.Columns {

		fieldList = append(fieldList, applyFieldSettings(&models.Field{
			Name: field.Name,
			Type: field.Type,
			Len:  field.Len,
		}, field.Settings))
	}

	return fieldList
}

func applyFieldSettings(field *models.Field, settings []*parseobj.FieldSetting) *models.Field {
	for _, setting := range settings {
		if setting.DefaultValue != nil {
			field.DefaultValue = *setting.DefaultValue
		}

		if setting.Note != nil {
			field.Note = *setting.Note
		}

		field.IsIncrement = field.IsIncrement || setting.Increment
		field.IsNotNull = field.IsNotNull || setting.NotNull
		field.IsPrimaryKey = field.IsPrimaryKey || setting.PrimaryKey
		field.IsUnique = field.IsUnique || setting.Unique
	}

	return field
}
