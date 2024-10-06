package generator

import (
	"github.com/vloldik/dbml-gen/internal/dbparse/models"
)

type IStructFromTableGenerator interface {
	ForGenerator(*DBMLGoGenerator) IStructFromTableGenerator
	CreateStructsFromTables([]*models.Table, *models.DBML) error
	Prepare(*models.DBML) error
	Finalize() error
}

type DBMLGoGenerator struct {
	From      *models.DBML
	StructGen IStructFromTableGenerator
	Module    string
	TagStyle  string
	OutputDIR string
}

func New(outputDIR, module, tagStyle string, structGen IStructFromTableGenerator) *DBMLGoGenerator {
	generator := &DBMLGoGenerator{
		OutputDIR: outputDIR,
		Module:    module,
		TagStyle:  tagStyle,
	}
	generator.StructGen = structGen.ForGenerator(generator)
	return generator
}

func (gen *DBMLGoGenerator) GenerateModels(parsed *models.DBML) error {
	if err := gen.StructGen.Prepare(parsed); err != nil {
		return err
	}
	if err := gen.StructGen.CreateStructsFromTables(parsed.Tables, parsed); err != nil {
		return err
	}
	return gen.StructGen.Finalize()
}
