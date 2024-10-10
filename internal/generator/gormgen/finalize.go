package gormgen

import (
	"github.com/iancoleman/strcase"
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

	if err := sg.createMigrator(); err != nil {
		return err
	}

	for _, generatedStruct := range sg.generatedStructs {
		if err := sg.createRepositrory(generatedStruct); err != nil {
			return err
		}
	}

	return nil
}
