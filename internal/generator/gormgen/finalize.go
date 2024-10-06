package gormgen

import (
	"sort"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/vloldik/dbml-gen/internal/utils/maputil"
)

const MigratorPackage = "migrates"
const gormPackage = "gorm.io/gorm"

// Finalize implements IStructFromTableGenerator.
func (sg *GORMStructGenerator) Finalize() error {
	for _, generated := range sg.generatedStructs {
		if err := saveFile(
			sg.parent.OutputDIR,
			generated.PackageNameToImport,
			strcase.ToSnake(generated.StructName),
			generated.File,
		); err != nil {
			return err
		}
	}

	migrateFile := jen.NewFile(MigratorPackage)
	migrateFile.Func().Id("MigrateAll").Call(jen.Id("db").Id("*").Qual(gormPackage, "DB")).Error().Block(
		jen.Return().Id("db").Dot("AutoMigrate").CallFunc(func(g *jen.Group) {
			values := maputil.Values(sg.generatedStructs)
			sort.Slice(values, func(i, j int) bool {
				return sg.RequirementsOrder(values[i]) >
					sg.RequirementsOrder(values[j])
			})
			for _, generated := range values {
				g.Id("&").Qual(generated.PackagePathToImport, generated.StructName).Block()
			}
		}),
	)

	return saveFile(sg.parent.OutputDIR, MigratorPackage, "migrate", migrateFile)
}
