package generator

import (
	"github.com/dave/jennifer/jen"
	"guthub.com/vloldik/dbml-gen/internal/utils/fileutil"
	"guthub.com/vloldik/dbml-gen/internal/utils/strutil"
)

const createFileMode = 677

func saveFile(path, packageName, filename string, file *jen.File) error {
	location := strutil.ConcantatePaths(path, packageName)
	if err := fileutil.EnsureFolderExists(location, createFileMode); err != nil {
		return err
	}

	return file.Save(strutil.ConcantatePaths(location, filename))
}
