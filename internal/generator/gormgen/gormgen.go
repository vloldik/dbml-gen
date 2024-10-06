package gormgen

import (
	"github.com/dave/jennifer/jen"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/generator"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

func NewStructGenerator() generator.IStructFromTableGenerator {
	return &GORMStructGenerator{
		generatedStructs:   make(map[uint32]*GeneratedStruct),
		structRequirements: make(map[uint32][]uint32),
	}
}

// CreateStructsFromTables implements generator.IStructFromTableGenerator.
func (sg *GORMStructGenerator) CreateStructsFromTables(tables []*models.Table, parsed *models.DBML) error {
	for _, table := range tables {

		sg.currentStruct = &GeneratedStruct{
			Source:     table,
			StructName: strutil.ToSingle(strutil.ToExportedGoName(table.Name.BaseName)),
			File:       jen.NewFile(table.Name.Namespace),

			PackagePathToImport: strutil.ConcantatePaths(sg.getBaseImportPath(), table.Name.Namespace),
			PackageNameToImport: table.Name.Namespace,
		}

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
		jen.Return(jen.Lit(sg.currentStruct.Source.Name.BaseName)),
	)
}
