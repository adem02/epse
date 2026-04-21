package project

import (
	"fmt"
	"path/filepath"

	"github.com/adem02/epse/internal/templates"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func CreateProjectFilesFromTemplate(basePath string, projectType typeutils.ProjectType, tmplData typeutils.TmplData) error {
	logutils.Logger{}.Info("📝 Création des fichiers du projet...")
	projectStructure, err := GetTemplatePathsByProjectType(projectType)
	if err != nil {
		panic(err)
	}

	for fileName, tmplFilePath := range projectStructure {
		fmt.Printf("🔍 File: %s → Template: %s\n", fileName, tmplFilePath) // ← DEBUG

		filePath := filepath.Join(basePath, fileName)
		if err := osutils.CreateFileFromTmpl(templates.FS, tmplFilePath, filePath, &tmplData); err != nil {
			return err
		}
	}

	return nil
}
