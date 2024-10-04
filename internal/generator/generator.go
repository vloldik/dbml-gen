package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
)

type IStructFromTableGenerator interface {
	CreateStruct(*models.DBML, *models.Table) (jen.Code, error)
}

type DBMLGoGenerator struct {
	From      *models.DBML
	StructGen IStructFromTableGenerator
}

func New() *DBMLGoGenerator {
	return &DBMLGoGenerator{
		StructGen: &StructGenerator{},
	}
}

// GenerateModels генерирует файлы моделей Go из таблиц БД.
func (gen *DBMLGoGenerator) GenerateModels(parsed *models.DBML, outputDir string, tagStyle string) error {
	for _, table := range parsed.Tables {
		file := jen.NewFile("models")

		code, err := gen.StructGen.CreateStruct(parsed, table)
		if err != nil {
			return err
		}

		file.Add(code)
		fileName := fmt.Sprintf("%s.go", strings.ToLower(table.Name))
		filePath := filepath.Join(outputDir, fileName)

		err = file.Save(filePath)
		if err != nil {
			return fmt.Errorf("ошибка при записи файла %s: %v", filePath, err)
		}
	}
	return nil
}
