package converts

import (
	"fmt"

	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
	"guthub.com/vloldik/dbml-gen/internal/utils/maputil"
)

type tableMap map[string]*models.Table

// Map [relation.Hash() <-> relation] for checks
type relationMap map[uint32]*models.Relationship

type ParseObjectToModelConverter struct {
	referenceList []*parseobj.StructureFullReference
	tableMap      tableMap
	relationMap   relationMap
}

func NewParseObjectToModelConverter() *ParseObjectToModelConverter {
	tables := make(tableMap)
	return &ParseObjectToModelConverter{
		tableMap:    tables,
		relationMap: make(relationMap),
	}
}

func (c *ParseObjectToModelConverter) processDBMLStructure(unknownStructure parseobj.DBMLStructure) error {
	switch structure := unknownStructure.(type) {
	case *parseobj.StructureFullReference:
		return c.processStructureReference(structure)
	case *parseobj.StructureTable:
		return c.processStructureTable(structure)
	case *parseobj.StructureEnum:
		//TODO
		return nil
	default:
		return fmt.Errorf("unknown structure type %T", structure)
	}
}

// This function converts parsed DBML to model and also checks logic
func (c *ParseObjectToModelConverter) ObjToModel(obj *parseobj.DBML) (*models.DBML, error) {
	for _, structure := range obj.Structures {
		if err := c.processDBMLStructure(structure); err != nil {
			return nil, err
		}
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

func (c *ParseObjectToModelConverter) processStructureTable(table *parseobj.StructureTable) error {
	fields, err := c.createFields(table)
	if err != nil {
		return err
	}
	tableModel := &models.Table{
		Name:   models.NewNamespacedName(table.Name.Namespace, table.Name.Name),
		Fields: fields,
		Alias:  table.As,
	}
	if err := c.applySettings(tableModel, table.Settings); err != nil {
		return err
	}
	c.tableMap[tableModel.Name.FullName()] = tableModel
	if tableModel.Alias != nil {
		c.tableMap[tableModel.Name.Namespace+"."+*tableModel.Alias] = tableModel
	}
	return nil
}

func (c *ParseObjectToModelConverter) createFields(table *parseobj.StructureTable) ([]*models.Field, error) {
	fieldList := make([]*models.Field, 0)
	for _, field := range table.Content.Columns {
		fieldModel := &models.Field{
			TableName: models.NewNamespacedName(table.Name.Namespace, table.Name.Name),
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
