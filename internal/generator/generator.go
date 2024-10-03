package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen" // Added for case conversion
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
	"guthub.com/vloldik/dbml-gen/internal/utils/genutil"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

type IStructFromTableGenerator interface {
	CreateStruct(*parseobj.DBML, *parseobj.Table) (*jen.Code, error)
}

type DBMLGoGenerator struct {
	From      *parseobj.DBML
	StructGen IStructFromTableGenerator
}

// GenerateModels generates Go model files from DB tables.
func GenerateModels(parsed *models.DBML, outputDir string, tagStyle string) error {
	for _, table := range parsed.Tables {
		file := generateModelContent(table, tagStyle)
		fileName := fmt.Sprintf("%s.go", strings.ToLower(table.Name))
		filePath := filepath.Join(outputDir, fileName)

		err := file.Save(filePath)
		if err != nil {
			return fmt.Errorf("error writing file %s: %v", filePath, err)
		}
	}
	return nil
}

// generateModelContent generates the content of a Go model file for a single table.
func generateModelContent(table *models.Table, _ string) *jen.File {
	file := jen.NewFile("models")

	structFields := make([]jen.Code, 0, len(table.Fields))

	for _, column := range table.Fields {
		goFieldName := strutil.ToExportedGoName(column.Name)
		goType := genutil.MapDBTypeToGoType(column.Type)

		field := jen.Id(goFieldName).Id(goType.GoString())

		structFields = append(structFields, field)
	}

	file.Type().Id(strutil.ToExportedGoName(table.Name)).Struct(structFields...)

	return file
}
