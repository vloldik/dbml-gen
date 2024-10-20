package gormgen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/vloldik/dbml-gen/internal/dbparse/models"
	"github.com/vloldik/dbml-gen/internal/generator"
	"github.com/vloldik/dbml-gen/internal/utils/genutil"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

type GORMStructGenerator struct {
	currentStruct    *GeneratedStruct
	generatedStructs map[uint32]*GeneratedStruct
	parent           *generator.DBMLGoGenerator
	structFields     generatedStructFields
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
	return nil
}

func (sg *GORMStructGenerator) createStruct(dbml *models.DBML, table *models.Table) error {
	sg.structFields = make(generatedStructFields, len(table.Fields))

	for i, field := range table.Fields {

		jenField := jen.Add()
		if field.Note != "" {
			jenField.Comment(field.NotePrepared()).Line()
		}
		jenField.Id(field.DisplayName())
		genutil.MapDBTypeToGoType(jenField, field.Type, field.IsNotNull)

		settings, err := field.CreateBasicGORMTags()
		if err != nil {
			return err
		}

		if len(field.Indexes) != 0 {
			for _, idx := range field.Indexes {
				settings = append(settings, sg.createIndexTags(idx, field)...)
			}
		}

		jenField.Tag(genutil.GormTagsFromList(settings...))

		sg.structFields[i] = &GeneratedField{
			Code: jenField,
			Name: field.DisplayName(),
		}

		for _, relation := range dbml.RelationsByFieldHash(field.Hash()) {
			sg.createFieldRelation(relation)
		}

	}
	structCode := jen.Add()
	if sg.currentStruct.Source.Note != "" {
		structCode.Comment(strutil.TryUnquote(sg.currentStruct.Source.Note)).Line()
	}
	structCode.Type().Id(sg.currentStruct.StructName).StructFunc(func(g *jen.Group) {
		sg.structFields.addAllTo(g)
	})

	sg.currentStruct.File.Add(structCode)
	return nil
}

func (sg *GORMStructGenerator) createFieldRelation(relation *models.Relationship) {
	var foreignKey, references string
	if relation.RelationType == models.OneToMany || relation.RelationType == models.ManyToOne {
		foreignKey = relation.ToField.DisplayName()
		references = relation.FromField.DisplayName()
	} else {
		references = relation.ToField.DisplayName()
		foreignKey = relation.FromField.DisplayName()
	}
	tags := []string{
		fmt.Sprintf("foreignKey:%s", foreignKey),
		fmt.Sprintf("References:%s", references),
	}
	if tag := GORMTagForRelationAction(relation); tag != "" {
		tags = append(tags, tag)
	}

	// True if we need import
	needSpecifyPackageName := relation.FromField.Table.PackageName() != relation.ToField.Table.PackageName()
	// True if we want to use []list
	isX_ToMany := relation.RelationType == models.ManyToMany || relation.RelationType == models.OneToMany
	qual := sg.getStructQualifier(relation.ToTable)
	createdFieldName := sg.createRelatedFieldName(relation.ToField, relation.ToTable, isX_ToMany)

	createdField := jen.Id(createdFieldName)
	if isX_ToMany {
		createdField.Index() //List []
	}

	if relation.RelationType == models.ManyToMany {
		tags = append(tags,
			fmt.Sprintf("many2many:%s", strutil.CreateManyToManyName(relation.FromTable.TableName.BaseName, relation.ToTable.TableName.BaseName)),
		)
	}

	createdField.Id("*") // Pointer
	if needSpecifyPackageName {
		createdField.Qual(qual, relation.ToTable.DisplayName())
	} else {
		createdField.Id(relation.ToTable.DisplayName())
	}

	createdField.Tag(genutil.GormTagsFromList(tags...))

	sg.structFields = append(sg.structFields, &GeneratedField{
		Code: createdField,
		Name: createdFieldName,
	})
}

func GORMTagForRelationAction(relation *models.Relationship) string {
	result := make([]string, 0)
	if relation.OnDelete == models.NoAction && relation.OnUpdate == models.NoAction {
		return ""
	}
	if relation.OnDelete != models.NoAction {
		result = append(result, fmt.Sprintf("OnDelete:%s", relation.OnDelete))
	}
	if relation.OnUpdate != models.NoAction {
		result = append(result, fmt.Sprintf("OnUpdate:%s", relation.OnUpdate))
	}

	return fmt.Sprintf("constraint:%s", strings.Join(result, ","))
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

func (sg *GORMStructGenerator) getStructQualifier(table *models.Table) string {
	qual := strutil.ConcantatePaths(sg.getBaseImportPath(), table.PackageName())
	return qual
}

func (sg *GORMStructGenerator) createRelatedFieldName(field *models.Field, table *models.Table, isX_toMany bool) string {
	relatedFieldName := field.DisplayName()
	relatedFieldName, found := strings.CutSuffix(relatedFieldName, "Id")
	if len(relatedFieldName) < 2 || !found {
		relatedFieldName = table.DisplayName()
	}
	if sg.structFields.hasName(relatedFieldName) {
		relatedFieldName = "Related" + relatedFieldName
	}
	if isX_toMany {
		relatedFieldName = strutil.ToPlural(relatedFieldName)
	}
	i := 0
	for uniqueName := relatedFieldName; sg.structFields.hasName(uniqueName); uniqueName = relatedFieldName + strconv.Itoa(i) {
		i++
	}
	if i != 0 {
		relatedFieldName = relatedFieldName + strconv.Itoa(i)
	}

	return relatedFieldName
}
