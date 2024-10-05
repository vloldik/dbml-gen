package converts

import (
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
	"guthub.com/vloldik/dbml-gen/internal/utils/maputil"
)

type tableMap map[string]*models.Table

// Map [relation.Hash() <-> relation] for checks
type relationMap map[uint32]*models.Relationship

type ParseObjectToModelConverter struct {
	parseDBML      *parseobj.DBML
	dbml           *models.DBML
	refsFromFields []*parseobj.FullReference
	tableMap       tableMap
	relationMap    relationMap
}

func NewParseObjectToModelConverter() *ParseObjectToModelConverter {
	tables := make(tableMap)
	return &ParseObjectToModelConverter{
		tableMap:    tables,
		relationMap: make(relationMap),
	}
}

// This function converts parsed DBML to model and also checks logic
func (c *ParseObjectToModelConverter) ObjToModel(obj *parseobj.DBML) (*models.DBML, error) {
	c.parseDBML = obj

	if err := c.fillTablesFromDBML(); err != nil {
		return nil, err
	}

	err := c.CreateRelations()
	if err != nil {
		return nil, err
	}

	return &models.DBML{
		Tables:    maputil.Values(c.tableMap),
		Relations: maputil.Values(c.relationMap),
	}, nil
}

func (c *ParseObjectToModelConverter) fillTablesFromDBML() error {
	for _, table := range c.parseDBML.Tables {
		fields, err := c.createFields(table)
		if err != nil {
			return err
		}
		tableModel := &models.Table{
			Name:   table.Name,
			Fields: fields,
		}
		if err := c.applySettings(tableModel, table.Settings); err != nil {
			return err
		}
		c.tableMap[table.Name] = tableModel
	}
	return nil
}

func (c *ParseObjectToModelConverter) createFields(table *parseobj.Table) ([]*models.Field, error) {
	fieldList := make([]*models.Field, 0)
	for _, field := range table.Content.Columns {
		fieldModel := &models.Field{
			TableName: table.Name,
			Name:      field.Name,
			Type:      field.Type,
			Len:       field.Len,
		}
		if err := c.applySettings(fieldModel, field.Settings); err != nil {
			return nil, err
		}
		fieldList = append(fieldList, fieldModel)
	}

	return fieldList, nil
}
