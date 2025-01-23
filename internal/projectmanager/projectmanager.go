package projectmanager

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/adem02/epse/internal/utils"
)

type ProjectManager struct {
	ProjectType utils.ProjectType
	ProjectName string
	Destination string
}

func New(projectType utils.ProjectType, projectName, projectDestination string) ProjectManager {
	return ProjectManager{
		ProjectType: projectType,
		ProjectName: projectName,
		Destination: projectDestination,
	}
}

func (pm ProjectManager) ProcessBaseStructureGeneration() error {
	basePath := fmt.Sprintf("%s%s", pm.Destination, pm.ProjectName)
	if err := createProjectStructure(basePath, pm.ProjectType); err != nil {
		return err
	}

	tmplData := struct {
		ProjectName string
	}{ProjectName: pm.ProjectName}
	if err := generateBaseFilesContentFromTemplate(basePath, pm.ProjectType, tmplData); err != nil {
		return err
	}

	return nil
}

func createProjectStructure(basePath string, projectType utils.ProjectType) error {
	_, err := os.Stat(basePath)
	if !os.IsNotExist(err) {
		errMessage := fmt.Sprintf(`
			❌ error generating project structure
			❌ folder already exists: %s
		`, basePath)

		return errors.New(errMessage)
	}

	projectStructure, err := utils.GetProjectStructureByType(projectType)
	if err != nil {
		return err
	}

	for _, directory := range projectStructure {
		if err := os.MkdirAll(fmt.Sprintf("%s/%s", basePath, directory), os.ModePerm); err != nil {
			if os.IsPermission(err) {
				fmt.Printf("❌ insufficient permissions to create the directory: %s\n", directory)
			} else {
				fmt.Printf("❌ error while creating the directory: %v\n", err)
			}
			return err
		}
	}

	return nil
}

func generateBaseFilesContentFromTemplate(basePath string, projectType utils.ProjectType, tmplData any) error {
	templatesAbsPath, err := utils.GetAbsolutePath()
	if err != nil {
		return err
	}

	tmplPathsByProjectType, err := utils.GetTemplatePathsByProjectType(projectType)
	if err != nil {
		return err
	}

	var outputFile *os.File
	for fileName, templatePath := range tmplPathsByProjectType {
		tmplPath := filepath.Join(templatesAbsPath, templatePath)

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			errorMessage := fmt.Errorf("❌ error parsing template: %s", tmplPath)
			return errorMessage
		}

		filePath := fmt.Sprintf("%s/%s", basePath, fileName)
		outputFile, err = os.Create(filePath)
		if err != nil {
			errorMessage := fmt.Errorf("❌ erreur lors de la création du fichier: %s", filePath)
			return errorMessage
		}

		err = tmpl.Execute(outputFile, tmplData)
		if err != nil {
			return err
		}
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			panic(err)
		}
	}(outputFile)

	return nil
}
