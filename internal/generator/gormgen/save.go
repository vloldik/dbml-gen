package gormgen

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/vloldik/dbml-gen/internal/utils/fileutil"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

const createFileMode = 677

func saveFile(basePath, packageName, filename string, file *jen.File) error {
	location := strutil.ConcantatePaths(basePath, packageName)
	if err := fileutil.EnsureFolderExists(location, createFileMode); err != nil {
		return err
	}

	if strings.HasSuffix(filename, "_test") {
		filename += "-model"
	}

	return file.Save(strutil.ConcantatePaths(location, filename+".go"))
}
