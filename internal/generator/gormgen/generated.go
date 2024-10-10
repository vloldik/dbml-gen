package gormgen

import (
	"hash/fnv"

	"github.com/dave/jennifer/jen"
	"github.com/vloldik/dbml-gen/internal/dbparse/models"
)

type GeneratedStruct struct {
	Source     *models.Table
	StructName string
	File       *jen.File

	RequiredStructHashes []uint32
	PackagePathToImport  string
	PackageNameToImport  string
}

func importHash(packageNameToImport, structName string) uint32 {
	hash := fnv.New32a()

	hash.Write([]byte(packageNameToImport))
	hash.Write([]byte(structName))

	return hash.Sum32()
}

func (g *GeneratedStruct) Hash() uint32 {
	return importHash(g.PackagePathToImport, g.StructName)
}

type GeneratedField struct {
	Name string
	Code *jen.Statement
}

type generatedStructFields []*GeneratedField

type IAdder interface {
	Add(...jen.Code) *jen.Statement
}

func (gfs generatedStructFields) addAllTo(what IAdder) {
	for _, gf := range gfs {
		what.Add(gf.Code)
	}
}

func (gfs generatedStructFields) hasName(name string) bool {
	for _, gf := range gfs {
		if gf == nil {
			continue
		}
		if gf.Name == name {
			return true
		}
	}

	return false
}
