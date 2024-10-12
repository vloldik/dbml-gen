package gormgen

import (
	"os"
	"text/template"

	"github.com/vloldik/dbml-gen/internal/utils/fileutil"
	"github.com/vloldik/dbml-gen/internal/utils/strutil"
)

func SaveTemplate(tpl, path, filename string, values any) error {
	parsedTemplate, err := template.ParseFiles(tpl)
	if err != nil {
		return err
	}

	if err := fileutil.EnsureFolderExists(path, 0777); err != nil {
		return err
	}

	file, err := os.OpenFile(strutil.ConcantatePaths(path, filename), os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	return parsedTemplate.Execute(file, values)
}
