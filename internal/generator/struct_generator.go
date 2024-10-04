package generator

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/utils/genutil"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

type StructGenerator struct{}

func NewStructGenerator() *StructGenerator {
	return &StructGenerator{}
}

func (sg *StructGenerator) CreateStruct(dbml *models.DBML, table *models.Table) (jen.Code, error) {
	structFields := make([]jen.Code, 0, len(table.Fields))

	for _, field := range table.Fields {
		goFieldName := strutil.ToExportedGoName(field.Name)
		jenField := jen.Id(goFieldName)
		genutil.MapDBTypeToGoType(jenField, field.Type)

		tags, err := genutil.CreateGORMTags(field)
		if err != nil {
			return nil, err
		}

		// Добавляем теги для индексов
		indexTags := sg.createIndexTags(table, field)
		if indexTags != "" {
			tags["gorm"] += " " + indexTags
		}

		jenField.Tag(tags)

		if field.Note != "" {
			str, err := strutil.UnquoteString(field.Note)
			if err != nil {
				return nil, err
			}
			jenField.Comment(str)
		}

		structFields = append(structFields, jenField)
	}

	// Добавляем поля для отношений
	for _, rel := range dbml.Relations {
		if rel.FromTable == table {
			relField := sg.createRelationField(rel)
			structFields = append(structFields, relField)
		}
	}

	structName := strutil.ToExportedGoName(table.Name)
	return jen.Type().Id(structName).Struct(structFields...), nil
}

func (sg *StructGenerator) createIndexTags(table *models.Table, field *models.Field) string {
	var indexTags []string

	for _, index := range table.Indexes {
		for _, indexField := range index.Fields {
			if indexField == field {
				tag := "index"
				if index.IsUnique {
					tag = "uniqueIndex"
				}
				// if index.Name != "" {
				// 	tag += ":" + index.Name
				// }
				indexTags = append(indexTags, tag)
			}
		}
	}

	if len(indexTags) > 0 {
		return fmt.Sprintf("gorm:\"%s\"", strings.Join(indexTags, ";"))
	}

	return ""
}

func (sg *StructGenerator) createRelationField(rel *models.Relationship) jen.Code {
	fieldName := strutil.ToExportedGoName(rel.ToTable.Name)
	if rel.RelationType == models.ManyToMany {
		fieldName += "s"
	}

	field := jen.Id(fieldName)

	switch rel.RelationType {
	case models.OneToOne:
		field.Op("*").Id(strutil.ToExportedGoName(rel.ToTable.Name))
	case models.ManyToMany:
		field.Index().Id(strutil.ToExportedGoName(rel.ToTable.Name))
	case models.ManyToOne:
		field.Op("*").Id(strutil.ToExportedGoName(rel.ToTable.Name))
	}

	tag := fmt.Sprintf("gorm:\"foreignKey:%s\"", rel.FromField.Name)
	field.Tag(map[string]string{"json": fieldName, "gorm": tag})

	comment := fmt.Sprintf("%s relationship", rel.RelationType.Name())
	field.Comment(comment)

	return field
}
