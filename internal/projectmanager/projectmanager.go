package projectmanager

import (
	"errors"
	"fmt"
	"os"

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

	formattedDependencies, err := utils.GetFormattedDependenciesByProjectType(pm.ProjectType)
	if err != nil {
		return err
	}

	tmplData := utils.TmplData{
		ProjectName:     pm.ProjectName,
		Dependencies:    formattedDependencies[utils.Dependencies],
		DevDependencies: formattedDependencies[utils.DevDependencies],
	}
	if err := generateBaseFilesContentFromTemplate(basePath, pm.ProjectType, tmplData); err != nil {
		return err
	}

	return nil
}

func createProjectStructure(basePath string, projectType utils.ProjectType) error {
	_, err := os.Stat(basePath)
	if !os.IsNotExist(err) {
		errMessage := fmt.Sprintf(`
	❌ error generating structure, project already exists: %s
		`, basePath)

		return errors.New(errMessage)
	}

	projectStructure, err := utils.GetProjectStructureByType(projectType)
	if err != nil {
		return err
	}

	for _, directory := range projectStructure {
		directoryPath := fmt.Sprintf("%s/%s", basePath, directory)
		if err := utils.CreateDirectory(directoryPath); err != nil {
			return err
		}
	}

	return nil
}

func generateBaseFilesContentFromTemplate(basePath string, projectType utils.ProjectType, tmplData utils.TmplData) error {
	tmplPathsByProjectType, err := utils.GetTemplatePathsByProjectType(projectType)
	if err != nil {
		return err
	}

	for fileName, tmplFilePath := range tmplPathsByProjectType {
		err = utils.CreateFileFromTmpl(tmplFilePath, fileName, basePath, tmplData)
		if err != nil {
			return err
		}
	}

	return nil
}
