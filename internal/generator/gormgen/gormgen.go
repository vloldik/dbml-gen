package gormgen

import (
	"github.com/dave/jennifer/jen"
	"github.com/vloldik/dbml-gen/internal/dbparse/models"
	"github.com/vloldik/dbml-gen/internal/generator"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

func NewStructGenerator() generator.IStructFromTableGenerator {
	return &GORMStructGenerator{
		generatedStructs: make(map[uint32]*GeneratedStruct),
	}
}

// CreateStructsFromTables implements generator.IStructFromTableGenerator.
func (sg *GORMStructGenerator) CreateStructsFromTables(tables []*models.Table, parsed *models.DBML) error {
	for _, table := range tables {

		sg.currentStruct = &GeneratedStruct{
			Source:     table,
			StructName: table.DisplayName(),
			File:       jen.NewFile(table.TableName.Namespace),

			PackagePathToImport: strutil.ConcantatePaths(sg.getBaseImportPath(), table.TableName.Namespace),
			PackageNameToImport: table.TableName.Namespace,
		}

		sg.currentStruct.File.PackageComment("Code generated from DBML. DO NOT EDIT")

		err := sg.createStruct(parsed, table)
		if err != nil {
			return err
		}

		sg.addTableNameFunc()

		sg.generatedStructs[sg.currentStruct.Hash()] = sg.currentStruct

	}

	return nil
}

func (sg *GORMStructGenerator) addTableNameFunc() {
	sg.currentStruct.File.Func().Params(
		jen.Id(sg.currentStruct.StructName),
	).Id("TableName").Params().String().Block(
		jen.Return(jen.Lit(sg.currentStruct.Source.TableName.BaseName)),
	)
}
