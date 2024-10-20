package converts

import (
	"fmt"

	"github.com/vloldik/dbml-gen/internal/dbparse/models"
	"github.com/vloldik/dbml-gen/internal/dbparse/parseobj"
)

func (c *ParseObjectToModelConverter) CreateRelations() error {
	for _, ref := range c.referenceList {
		if err := c.createRelationsFromStructureReference(ref); err != nil {
			return err
		}
	}

	return nil
}

func (c *ParseObjectToModelConverter) processStructureReference(ref *parseobj.StructureFullReference) error {
	c.referenceList = append(c.referenceList, ref)
	return nil
}

func (c *ParseObjectToModelConverter) createRelationsFromStructureReference(ref *parseobj.StructureFullReference) error {
	field, err := normalizeRef(ref.Field)
	if err != nil {
		return err
	}
	refToField, err := normalizeRef(ref.ReferenceToField)
	if err != nil {
		return err
	}
	if relation, err := c.createRelation(
		field.NameParts[0],
		field.NameParts[1],
		field.NameParts[2],
		refToField.NameParts[0],
		refToField.NameParts[1],
		refToField.NameParts[2],
		ref,
	); err == nil {
		if err := c.addRelation(relation); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func (c *ParseObjectToModelConverter) createRelation(
	fromTableName, fromTableNamespace, fromColumnName string,
	toTableName, toTableNamespace, toColumnName string,
	fromObj *parseobj.StructureFullReference) (*models.Relationship, error) {

	createdRelationType, err := models.RelationTypeFromParsed(fromObj.Type)
	fromTableNameModel := models.NewNamespacedNameSafe(fromTableName, fromTableNamespace)
	toTableNameModel := models.NewNamespacedNameSafe(toTableName, toTableNamespace)

	if err != nil {
		return nil, err
	}

	fromTable, ok := c.tableMap[fromTableNameModel.FullName()]
	if !ok {
		return nil, fmt.Errorf("table %s did not found", fromTableNameModel.FullName())
	}
	fromField, err := fromTable.GetFieldByName(fromColumnName)
	if err != nil {
		return nil, err
	}

	toTable, ok := c.tableMap[toTableNameModel.FullName()]
	if !ok {
		return nil, fmt.Errorf("table %s did not found", toTableNameModel.FullName())
	}
	toField, err := toTable.GetFieldByName(toColumnName)
	if err != nil {
		return nil, err
	}

	relation := &models.Relationship{
		RelationType: createdRelationType,
		OnUpdate:     models.NoAction,
		OnDelete:     models.NoAction,

		FromTable: fromTable,
		FromField: fromField,

		ToTable: toTable,
		ToField: toField,
	}
	return relation, c.applySettings(relation, fromObj.Settings)
}

func (c *ParseObjectToModelConverter) addRelation(relation *models.Relationship) error {
	relHash := relation.Hash()
	if _, ok := c.relationMap[relHash]; ok {
		return fmt.Errorf(
			"duplicate relations %s.%s - %s.%s",
			relation.FromTable.TableName, relation.FromField.DBName,
			relation.ToTable.TableName, relation.ToField.DBName,
		)
	}

	c.relationMap[relHash] = relation
	return nil
}

func normalizeRef(ref *parseobj.ReferenceField) (*parseobj.ReferenceField, error) {
	partsLen := len(ref.NameParts)

	if partsLen == 2 {
		partsLen++
		ref.NameParts = append([]string{models.DefaultNamespace}, ref.NameParts...)
	}

	if partsLen != 3 {
		return nil, fmt.Errorf("reference field name parts count is invalid: %d", partsLen)
	}

	return ref, nil
}
