package converts

import (
	"fmt"

	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
)

func (c *ParseObjectToModelConverter) CreateRelations() error {
	if err := c.createRelationsFromFullRefList(c.parseDBML.Refs); err != nil {
		return err
	}
	if err := c.createRelationsFromFullRefList(c.refsFromFields); err != nil {
		return err
	}
	return nil
}

func (c *ParseObjectToModelConverter) createRelationsFromFullRefList(reflist []*parseobj.FullReference) error {
	for _, ref := range reflist {
		if relation, err := c.createRelation(
			ref.Field.Table,
			ref.Field.Column,
			ref.ReferenceToField.Table,
			ref.ReferenceToField.Column,
			ref.Type,
		); err == nil {
			if err := c.addRelation(relation); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (c *ParseObjectToModelConverter) createRelation(
	fromTableName, fromColumnName,
	toTableName, toColumnName string,
	relationshipType *parseobj.RelationshipType) (*models.Relationship, error) {

	createdRelationType, err := models.RelationTypeFromParsed(relationshipType)

	if err != nil {
		return nil, err
	}

	fromTable, ok := c.tableMap[fromTableName]
	if !ok {
		return nil, fmt.Errorf("table %s did not found", fromTableName)
	}
	fromField, err := fromTable.GetFieldByName(fromColumnName)
	if err != nil {
		return nil, err
	}

	toTable, ok := c.tableMap[toTableName]
	if !ok {
		return nil, fmt.Errorf("table %s did not found", toTableName)
	}
	toField, err := toTable.GetFieldByName(toColumnName)
	if err != nil {
		return nil, err
	}

	return &models.Relationship{
		RelationType: createdRelationType,

		FromTable: fromTable,
		FromField: fromField,

		ToTable: toTable,
		ToField: toField,
	}, nil
}

func (c *ParseObjectToModelConverter) addRelation(relation *models.Relationship) error {
	relHash := relation.Hash()
	if _, ok := c.relationMap[relHash]; ok {
		return fmt.Errorf(
			"duplicate relations %s.%s - %s.%s",
			relation.FromTable.Name, relation.FromField.Name,
			relation.ToTable.Name, relation.ToField.Name,
		)
	}

	c.relationMap[relHash] = relation
	return nil
}
