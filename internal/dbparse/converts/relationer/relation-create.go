package relationer

import (
	"fmt"

	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
	"guthub.com/vloldik/dbml-gen/internal/utils/maputil"
)

// Map [relation.Hash() <-> relation] for checks
type relationMap map[uint32]*models.Relationship
type tableMap map[string]*models.Table

type RelationCreator struct {
	relationMap relationMap
	tableMap    tableMap
}

func NewRelationCreator(tableMap map[string]*models.Table) *RelationCreator {
	return &RelationCreator{
		tableMap:    tableMap,
		relationMap: make(relationMap),
	}
}

func (r *RelationCreator) CreateRelations(dbml *parseobj.DBML) ([]*models.Relationship, error) {
	for _, table := range dbml.Tables {
		if err := r.createRelationsFromTableFields(table); err != nil {
			return nil, err
		}
	}

	for _, ref := range dbml.Refs {
		if relation, err := r.createRelation(
			ref.Field.Table,
			ref.Field.Column,
			ref.ReferenceToField.Table,
			ref.ReferenceToField.Column,
			ref.Type,
		); err == nil {
			if err := r.addRelation(relation); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return maputil.Values(r.relationMap), nil
}

func (r *RelationCreator) createRelationsFromTableFields(table *parseobj.Table) error {
	if table.Content == nil {
		return nil
	}
	// I hate nested for(((
	for _, column := range table.Content.Columns {
		for _, setting := range column.Settings {
			if setting.Reference == nil {
				continue
			}
			if relation, err := r.createRelation(
				table.Name, column.Name,
				setting.Reference.Table,
				setting.Reference.Column,
				setting.Reference.Type,
			); err == nil {
				if err := r.addRelation(relation); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}

func (r *RelationCreator) createRelation(
	fromTableName, fromColumnName,
	toTableName, toColumnName string,
	relationshipType *parseobj.RelationshipType) (*models.Relationship, error) {

	var createdRelationType models.RelationType

	if relationshipType.OneToOne {
		createdRelationType = models.OneToOne
	} else if relationshipType.ManyToOne {
		createdRelationType = models.ManyToOne
	} else if relationshipType.ManyToMany {
		createdRelationType = models.ManyToMany
	} else {
		// I hope this never happen
		return nil, fmt.Errorf("unknown relation type")
	}

	fromTable, ok := r.tableMap[fromTableName]
	if !ok {
		return nil, fmt.Errorf("table %s did not found", fromTableName)
	}
	fromField, err := fromTable.GetFieldByName(fromColumnName)
	if err != nil {
		return nil, err
	}

	toTable, ok := r.tableMap[toTableName]
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

func (r *RelationCreator) addRelation(relation *models.Relationship) error {
	relHash := relation.Hash()
	if _, ok := r.relationMap[relHash]; ok {
		return fmt.Errorf(
			"duplicate relations %s.%s - %s.%s",
			relation.FromTable.Name, relation.FromField.Name,
			relation.ToTable.Name, relation.ToField.Name,
		)
	}

	relation.FromField.ReferencesTo = append(relation.FromField.ReferencesTo, &models.Reference{
		RelationType: relation.RelationType,
		SecondField:  relation.ToField,
		SecondTable:  relation.ToTable,
	})

	relation.ToField.ReferencedBy = append(relation.FromField.ReferencedBy, &models.Reference{
		RelationType: relation.RelationType,
		SecondField:  relation.FromField,
		SecondTable:  relation.FromTable,
	})

	r.relationMap[relHash] = relation
	return nil
}
