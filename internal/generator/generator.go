package generator

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"guthub.com/vloldik/dbml-gen/internal/dbparse/models"
)

type IStructFromTableGenerator interface {
	CreateStruct(string, *models.DBML, *models.Table) (jen.Code, error)
}

type DBMLGoGenerator struct {
	From      *models.DBML
	StructGen IStructFromTableGenerator
	file      *jen.File
	module    string
	tagStyle  string
	outputDIR string
}

func New(outputDIR, module string, tagStyle string) *DBMLGoGenerator {
	return &DBMLGoGenerator{
		outputDIR: outputDIR,
		module:    module,
		tagStyle:  tagStyle,
	}
}

// GenerateModels генерирует файлы моделей Go из таблиц БД.
func (gen *DBMLGoGenerator) GenerateModels(parsed *models.DBML) error {
	for _, table := range parsed.Tables {
		gen.file = jen.NewFile(table.Name.Namespace)
		generator := NewStructGenerator(gen)

		err := generator.CreateStruct(parsed, table)
		if err != nil {
			return err
		}

		fileName := fmt.Sprintf("%s.go", strings.ToLower(table.Name.BaseName))

		err = saveFile(gen.outputDIR, table.Name.Namespace, fileName, gen.file)
		if err != nil {
			return fmt.Errorf("ошибка при записи файла %v", err)
		}
	}
	return nil
}
