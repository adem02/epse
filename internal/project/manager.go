package project

import (
	"fmt"
	"path/filepath"

	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type ProjectManager struct {
	ProjectType typeutils.ProjectType
	ProjectName string
	Destination string
}

func (pm *ProjectManager) Generate() error {
	projectPath := filepath.Join(pm.Destination, pm.ProjectName)
	// log project path
	logutils.Logger{}.Info(fmt.Sprintf("📂 Creating project %s...", projectPath))
	if err := CreateProjectStructureByType(projectPath, pm.ProjectType); err != nil {
		return err
	}

	if err := config.GenerateNewConfigFile(pm.ProjectType, pm.ProjectName, projectPath); err != nil {
		return err
	}

	formattedDependencies, err := GetFormattedDependenciesByProjectType(pm.ProjectType)
	if err != nil {
		return err
	}

	tmplData := typeutils.TmplData{
		ProjectName:     pm.ProjectName,
		Dependencies:    formattedDependencies[typeutils.Dependencies],
		DevDependencies: formattedDependencies[typeutils.DevDependencies],
	}

	if err := CreateProjectFilesFromTemplate(projectPath, pm.ProjectType, tmplData); err != nil {
		return err
	}

	displayEndingMessage(pm.ProjectName, pm.Destination, pm.ProjectType)

	return nil
}

func New(projectType typeutils.ProjectType, projectName, projectDestination string) *ProjectManager {
	return &ProjectManager{
		ProjectType: projectType,
		ProjectName: projectName,
		Destination: projectDestination,
	}
}

func displayEndingMessage(projectName, destination string, projectType typeutils.ProjectType) {
	logutils.Logger{}.Success("✅ Generation completed successfully!")

	logutils.Logger{}.Section("📂 Generated project", projectName)
	logutils.Logger{}.Section("📍 Location", destination)
	logutils.Logger{}.Section("🏗️ Project type", projectType)
	fmt.Println()

	logutils.Logger{}.Warning("🚀 Install dependencies")
	logutils.Logger{}.Info("   npm install\n")

	logutils.Logger{}.Warning("🚀 Run the project")
	logutils.Logger{}.Info("   npm run dev\n")
}
