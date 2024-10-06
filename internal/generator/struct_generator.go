package generator

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/utils/genutil"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

type StructGenerator struct {
	parent       *DBMLGoGenerator
	structFields []jen.Code
}

func NewStructGenerator(parent *DBMLGoGenerator) *StructGenerator {
	return &StructGenerator{
		parent: parent,
	}
}

func (sg *StructGenerator) CreateStruct(dbml *models.DBML, table *models.Table) error {
	sg.structFields = make([]jen.Code, len(table.Fields))

	for i, field := range table.Fields {
		goFieldName := strutil.ToExportedGoName(field.Name)
		jenField := jen.Id(goFieldName)
		genutil.MapDBTypeToGoType(jenField, field.Type)

		settings, err := genutil.CreateBasicGORMTags(field)
		if err != nil {
			return err
		}

		jenField.Tag(map[string]string{"gorm": strings.Join(settings, ";")})

		sg.structFields[i] = jenField

		if field.Relations == nil {
			continue
		}
		for _, relation := range field.Relations {
			sg.createFieldRelation(field, relation)
		}
	}

	structName := strutil.ToExportedGoName(table.Name.BaseName)
	sg.parent.file.Add(jen.Type().Id(structName).Struct(sg.structFields...))
	return nil
}

func (sg *StructGenerator) createFieldRelation(field *models.Field, relation *models.FieldRelation) {
	needSpecifyPackageName := false

	if relation.SecondTable.Namespace != field.TableName.Namespace {
		needSpecifyPackageName = true
	}
	relatedTypeName := strutil.ToExportedGoName(relation.SecondTable.BaseName)

	createdField := jen.Id(relatedTypeName)

	if relation.RelationType == models.ManyToMany || relation.RelationType == models.OneToMany {
		createdField.Index()
	}

	createdField.Id("*")
	if needSpecifyPackageName {
		importPackageName := strutil.ConcantatePaths(sg.parent.module, sg.parent.outputDIR, relation.SecondTable.Namespace)
		createdField.Qual(importPackageName, relatedTypeName)
	} else {
		createdField.Id(relatedTypeName)
	}

	sg.structFields = append(sg.structFields, createdField)
}
