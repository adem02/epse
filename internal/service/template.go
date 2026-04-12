package service

import (
	"fmt"
	"path/filepath"

	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func CreateServiceFileFromTmpl(projectType typeutils.ProjectType, name, destPath string) (bool, error) {
	if osutils.FileOrDirectoryExists(destPath) {
		logutils.Logger{}.Warning(fmt.Sprintf("⚠️ File already exists, skipping: %s", destPath))
		return false, nil
	}

	if err := osutils.CreateDirectory(filepath.Dir(destPath)); err != nil {
		return false, err
	}

	tmplPath := GetServiceTemplatePath(projectType)
	data := &struct{ ServiceName string }{
		ServiceName: Capitalize(name),
	}

	return true, osutils.CreateFileFromTmpl(tmplPath, destPath, data)
}
