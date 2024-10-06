package gormgen

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/generator"
	"guthub.com/vloldik/dbml-gen/internal/utils/genutil"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

type GORMStructGenerator struct {
	currentStruct      *GeneratedStruct
	generatedStructs   map[uint32]*GeneratedStruct
	structRequirements map[uint32][]uint32
	parent             *generator.DBMLGoGenerator
	structFields       []jen.Code
}

// ForGenerator implements IStructFromTableGenerator.
func (sg *GORMStructGenerator) ForGenerator(parent *generator.DBMLGoGenerator) generator.IStructFromTableGenerator {
	sg.parent = parent
	return sg
}

type IStructFromTableGenerator interface {
	Prepare(*models.DBML) error
	ForGenerator(*generator.DBMLGoGenerator) *GORMStructGenerator
	CreateStruct(*models.DBML, *models.Table) (string, error)
	Finalize() error
}

func (sg *GORMStructGenerator) Prepare(dbml *models.DBML) error {
	for _, relation := range dbml.Relations {
		requiredName, requiredQual := sg.structNameAndQual(relation.ToField.TableName)
		referencedName, referencedQual := sg.structNameAndQual(relation.FromField.TableName)
		requiredHash := importHash(requiredQual, requiredName)
		referencedHash := importHash(referencedQual, referencedName)

		sg.structRequirements[referencedHash] = append(sg.structRequirements[referencedHash], requiredHash)

	}

	return nil
}

func (sg *GORMStructGenerator) createStruct(dbml *models.DBML, table *models.Table) error {
	sg.structFields = make([]jen.Code, len(table.Fields))

	for i, field := range table.Fields {
		goFieldName := strutil.ToExportedGoName(field.Name)
		jenField := jen.Id(goFieldName)
		genutil.MapDBTypeToGoType(jenField, field.Type)

		settings, err := genutil.CreateBasicGORMTags(field)
		if err != nil {
			return err
		}

		if len(field.Indexes) != 0 {
			for _, idx := range field.Indexes {
				settings = append(settings, sg.createIndexTags(idx, field)...)
			}
		}

		jenField.Tag(genutil.GormTagsFromList(settings...))

		sg.structFields[i] = jenField

		if field.Relations == nil {
			continue
		}
		for _, relation := range field.Relations {
			sg.createFieldRelation(field, relation)
		}

	}

	sg.currentStruct.File.Add(jen.Type().Id(sg.currentStruct.StructName).Struct(sg.structFields...))
	return nil
}

func (sg *GORMStructGenerator) createFieldRelation(field *models.Field, relation *models.FieldRelation) {
	tags := []string{
		fmt.Sprintf("foreignKey:%s", field.Name),
		fmt.Sprintf("References:%s", relation.SecondField),
	}
	// True if we need import
	needSpecifyPackageName := false
	// True if we want to use []list
	isX_ToMany := relation.RelationType == models.ManyToMany || relation.RelationType == models.OneToMany

	if relation.SecondTable.Namespace != field.TableName.Namespace {
		needSpecifyPackageName = true
	}

	relatedTypeName, qual := sg.structNameAndQual(relation.SecondTable)
	relatedFieldName := relatedTypeName
	if isX_ToMany {
		relatedFieldName = strutil.ToPlural(relatedFieldName)
	}

	createdField := jen.Id(relatedFieldName)
	if isX_ToMany {
		createdField.Index() //List []
	}

	if relation.RelationType == models.ManyToMany {
		tags = append(tags,
			fmt.Sprintf("many2many:%s", strutil.CreateManyToManyName(field.TableName.BaseName, relation.SecondTable.BaseName)),
		)
	}

	createdField.Id("*") // Pointer
	if needSpecifyPackageName {
		createdField.Qual(qual, relatedTypeName)
	} else {
		createdField.Id(relatedTypeName)
	}

	createdField.Tag(genutil.GormTagsFromList(tags...))

	requirements, ok := sg.structRequirements[sg.currentStruct.Hash()]
	if ok {
		sg.currentStruct.RequiredStructHashes = requirements
	}
	sg.structFields = append(sg.structFields, createdField)
}

func (sg *GORMStructGenerator) createIndexTags(index *models.Index, _ *models.Field) (tags []string) {
	firstPart := "index:" + index.GetName()

	indexParts := []string{firstPart}

	if index.IsUnique {
		indexParts = append(indexParts, "unique")
	}

	if index.Type != "" {
		indexParts = append(indexParts, "type:"+index.Type)
	}

	if len(index.Exprs) != 0 {
		exprs := []string{}
		for _, expr := range index.Exprs {
			expr, _ = strutil.RemoveQuotes(expr, "`")
			exprs = append(exprs, expr)
		}
		indexParts = append(indexParts, fmt.Sprintf("expression:%s", strings.Join(exprs, `\\,`)))
	}

	if index.IsPrimaryKey {
		tags = append(tags, "primaryKey")
	}

	return append(tags, strings.Join(indexParts, ","))
}

func (sg *GORMStructGenerator) getBaseImportPath() string {
	return strutil.NormalizePath(strutil.ConcantatePaths(sg.parent.Module))
}

func (sg *GORMStructGenerator) structNameAndQual(tableName *models.NamespacedName) (string, string) {
	structName := strutil.ToSingle(strutil.ToExportedGoName(tableName.BaseName))
	qual := strutil.ConcantatePaths(sg.getBaseImportPath(), tableName.Namespace)
	return structName, qual
}
