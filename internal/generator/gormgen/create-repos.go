package gormgen

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/vloldik/dbml-gen/internal/dbparse/models"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

const (
	model                = "model"
	optionsPackage       = "opt"
	codeGeneratedComment = "Code generated from DBML. DO NOT EDIT"
)

func (sg *GORMStructGenerator) CreateRepositories(dbml *models.DBML) error {
	for _, table := range dbml.Tables {
		if err := sg.createRepositrory(table); err != nil {
			return err
		}
	}
	return nil
}

type templateOptions struct {
	Package     string
	GORMPackage string
}

func (sg *GORMStructGenerator) createRepositrory(generatedStruct *models.Table) error {
	if err := SaveTemplate(
		"./template/apply-option.tpl",
		strutil.ConcantatePaths(sg.parent.OutputDIR, optionsPackage),
		"options.go",
		templateOptions{
			Package:     optionsPackage,
			GORMPackage: gormPackage,
		},
	); err != nil {
		return err
	}
	repoName := generatedStruct.DisplayName() + "Repository"
	qual := sg.getStructQualifier(generatedStruct)

	file := jen.NewFile("repos")
	file.PackageComment(codeGeneratedComment)
	file.ImportName("gorm.io/gorm", "gorm")

	file.Type().Id(repoName).StructFunc(func(g *jen.Group) {
		g.Id("db").Id("*").Qual("gorm.io/gorm", "DB")
	})

	file.Func().Id("New" + repoName).Params(jen.Id("db").Id("*").Qual("gorm.io/gorm", "DB")).Id("*").Id(repoName).BlockFunc(func(g *jen.Group) {
		g.Return(jen.Id("&").Id(repoName).Values(jen.Dict{
			jen.Id("db"): jen.Id("db"),
		}))
	})

	funcCreator := &funcCreator{
		file:           file,
		repoClassName:  repoName,
		packageName:    qual,
		baseImportPath: sg.getBaseImportPath(),
		className:      generatedStruct.DisplayName(),
	}

	for _, field := range generatedStruct.Fields {
		if !field.IsPrimaryKey {
			continue
		}

		funcCreator.GenFuncGetBy(field)
		funcCreator.GenFuncDelete(field)
	}

	funcCreator.GenFuncCreate()
	funcCreator.GenFuncList()
	funcCreator.GenFuncUpdate()
	funcCreator.GenFuncTotalCount()

	funcCreator.GenFuncGetDB()

	return saveFile(sg.parent.OutputDIR, "repos", strcase.ToSnake(generatedStruct.DisplayName())+"_repo", file)
}

func writeClassFunc(writeTo *jen.File, classname, funcName string) *jen.Statement {
	return writeTo.Func().Params(jen.Id("r").Id("*").Id(classname)).Id(funcName)
}

type funcCreator struct {
	file           *jen.File
	repoClassName  string
	baseImportPath string
	packageName    string
	className      string
}

func (fc *funcCreator) GenBlockApplyOptions() jen.Code {
	optsPackageQual := strutil.ConcantatePaths(fc.baseImportPath, optionsPackage)
	return jen.Qual(optsPackageQual, "ApplyOptions").CallFunc(func(g *jen.Group) {
		g.Id("r").Dot("db").Dot("WithContext").Call(jen.Id("ctx"))
		g.Id("opts").Op("...")
	})
}

func (fc *funcCreator) GenFuncGetDB() jen.Code {
	statement := writeClassFunc(fc.file, fc.repoClassName, "GetDB")
	return statement.Params().Id("*").Qual(gormPackage, "DB").BlockFunc(func(g *jen.Group) {
		g.Return().Id("r").Dot("db")
	})
}

func (fc *funcCreator) optsParam() jen.Code {
	return jen.Id("opts").Op("...").Any()
}

func (fc *funcCreator) GenFuncGetBy(field *models.Field) jen.Code {
	statement := writeClassFunc(fc.file, fc.repoClassName, "GetBy"+field.DisplayName())
	fieldNameLower := strutil.ToNotExported(field.DisplayName())
	lowercaseClass := strutil.ToNotExported(fc.className)
	return statement.Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id(fieldNameLower).Any(),
		fc.optsParam(),
	).Params(
		jen.Op("*").Qual(fc.packageName, fc.className), jen.Id("error"),
	).BlockFunc(func(g *jen.Group) {
		g.Var().Id(lowercaseClass).Qual(fc.packageName, fc.className)
		g.Id("result").Op(":=").Add(fc.GenBlockApplyOptions()).
			Dot("First").Call(
			jen.Id("&").Id(lowercaseClass), jen.Lit(fmt.Sprintf("%s = ?", field.DBName)), jen.Id(fieldNameLower),
		)

		g.Add(genReturnErrorIfNotNil())

		g.Return().List(jen.Id("&").Id(lowercaseClass), jen.Id("nil"))
	})
}

func (fc *funcCreator) GenFuncCreate() *jen.Statement {
	statement := writeClassFunc(fc.file, fc.repoClassName, "Create")
	return statement.Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id(model).Qual(fc.packageName, fc.className),
		fc.optsParam(),
	).Params(
		jen.Op("*").Qual(fc.packageName, fc.className), jen.Id("error"),
	).BlockFunc(func(g *jen.Group) {
		g.Id("result").Op(":=").Add(fc.GenBlockApplyOptions()).
			Dot("Create").Call(
			jen.Id("&").Id(model),
		)

		g.Add(genReturnErrorIfNotNil())

		g.Return().List(jen.Id("&").Id(model), jen.Id("nil"))
	})
}

func (fc *funcCreator) GenFuncList() *jen.Statement {
	statement := writeClassFunc(fc.file, fc.repoClassName, "List")
	return statement.Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("limit").Int(),
		jen.Id("offset").Int(),
		fc.optsParam(),
	).Params(
		jen.Index().Op("*").Qual(fc.packageName, fc.className), jen.Id("error"),
	).BlockFunc(func(g *jen.Group) {
		g.Var().Id("list").Index().Id("*").Qual(fc.packageName, fc.className)
		g.Id("result").Op(":=").Add(fc.GenBlockApplyOptions()).
			Dot("Limit").Call(jen.Id("limit")).
			Dot("Offset").Call(jen.Id("offset")).
			Dot("Find").Call(jen.Id("&").Id("list"))

		g.Add(genReturnErrorIfNotNil())

		g.Return().List(jen.Id("list"), jen.Id("nil"))
	})
}

func (fc *funcCreator) GenFuncDelete(field *models.Field) *jen.Statement {
	statement := writeClassFunc(fc.file, fc.repoClassName, "DeleteBy"+field.DisplayName())
	fieldNameLower := strutil.ToNotExported(field.DisplayName())
	lowercaseClass := strutil.ToNotExported(fc.className)
	return statement.Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id(fieldNameLower).Id("any"),
		fc.optsParam(),
	).Params(
		jen.Id("error"),
	).BlockFunc(func(g *jen.Group) {
		g.Var().Id(lowercaseClass).Qual(fc.packageName, fc.className)
		g.Id("result").Op(":=").Add(fc.GenBlockApplyOptions()).
			Dot("Delete").Call(
			jen.Id("&").Id(lowercaseClass), jen.Lit(fmt.Sprintf("%s = ?", field.DBName)), jen.Id(fieldNameLower),
		)

		g.Add(genReturnErrorIfNotNil(""))

		g.Return().Id("nil")
	})
}

func (fc *funcCreator) GenFuncUpdate() *jen.Statement {
	statement := writeClassFunc(fc.file, fc.repoClassName, "Update")
	return statement.Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id(model).Qual(fc.packageName, fc.className),
		fc.optsParam(),
	).Params(
		jen.Op("*").Qual(fc.packageName, fc.className), jen.Id("error"),
	).BlockFunc(func(g *jen.Group) {
		g.Id("result").Op(":=").Add(fc.GenBlockApplyOptions()).
			Dot("Updates").Call(jen.Id("&").Id(model))

		g.Add(genReturnErrorIfNotNil())

		g.Return().List(jen.Id("&").Id(model), jen.Id("nil"))
	})
}

func (fc *funcCreator) GenFuncTotalCount() *jen.Statement {
	statement := writeClassFunc(fc.file, fc.repoClassName, "TotalCount")
	return statement.Params(
		jen.Id("ctx").Qual("context", "Context"),
	).Params(
		jen.Id("int64"), jen.Id("error"),
	).BlockFunc(func(g *jen.Group) {
		g.Var().Id("count").Int64()
		g.Id("result").Op(":=").Id("r").Dot("db").
			Dot("WithContext").Call(jen.Id("ctx")).
			Dot("Model").Call(jen.Op("&").Qual(fc.packageName, fc.className).Values()).
			Dot("Count").Call(jen.Op("&").Id("count"))

		g.Add(genReturnErrorIfNotNil("-1"))

		g.Return().List(jen.Id("count"), jen.Id("nil"))
	})
}

func genReturnErrorIfNotNil(defaultReturn ...string) *jen.Statement {
	var returnValue string
	if len(defaultReturn) == 0 {
		returnValue = "nil"
	} else {
		returnValue = defaultReturn[0]
	}
	statement := jen.Return()
	if returnValue == "" {
		statement.List(jen.Id("result").Dot("Error"))
	} else {
		statement.List(jen.Id(returnValue), jen.Id("result").Dot("Error"))
	}
	return jen.If(jen.Id("result").Dot("Error").Op("!=").Id("nil")).Block(statement)
}
