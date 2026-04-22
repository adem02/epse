package repository

import (
	"fmt"
	"path/filepath"

	"github.com/adem02/epse/internal/templates"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func CreateRepositoryFileFromTmpl(projectType typeutils.ProjectType, names RepositoryNames, templateName, destPath string) (bool, error) {
	if osutils.FileOrDirectoryExists(destPath) {
		logutils.Logger{}.Warning(fmt.Sprintf("⚠️ File already exists, skipping: %s", destPath))
		return false, nil
	}

	if err := osutils.CreateDirectory(filepath.Dir(destPath)); err != nil {
		return false, err
	}

	tmplPath := GetRepositoryTemplatePath(projectType, templateName)
	data := &struct {
		RepositoryName      string
		InterfaceImportPath string
	}{
		RepositoryName:      names.CleanName,
		InterfaceImportPath: names.InterfaceImportPath,
	}

	return true, osutils.CreateFileFromTmpl(templates.FS, tmplPath, destPath, data)
}
