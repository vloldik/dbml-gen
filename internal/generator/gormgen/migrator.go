package gormgen

import (
	"github.com/dave/jennifer/jen"
	"github.com/vloldik/dbml-gen/internal/utils/maputil"
)

func (sg *GORMStructGenerator) createMigrator() error {
	migrateFile := jen.NewFile(MigratorPackage)
	migrateFile.Func().Id("MigrateAll").Call(jen.Id("db").Id("*").Qual(gormPackage, "DB")).Error().Block(
		jen.Return().Id("db").Dot("AutoMigrate").CallFunc(func(g *jen.Group) {
			for _, generated := range maputil.Values(sg.generatedStructs) {
				g.Id("&").Qual(generated.PackagePathToImport, generated.StructName).Block()
			}
		}),
	)

	return saveFile(sg.parent.OutputDIR, MigratorPackage, "migrate", migrateFile)
}
