package generator

import (
	"errors"
	"fmt"
	"github.com/adem02/epse/internal/projectmanager"
	"github.com/adem02/epse/internal/utils"
	"strings"
)

type Generator struct {
	ProjectType utils.ProjectType
	ProjectName string
	Destination string
}

func generateLiteStructure(projectName, destination string) error {
	projectManager := projectmanager.New(utils.LiteProjectType, projectName, destination)
	if err := projectManager.ProcessBaseStructureGeneration(); err != nil {
		return err
	}

	displayEndingMessage(projectName, destination, utils.LiteProjectType)

	return nil
}

func generateCleanStructure(projectName, destination string) error {
	projectManager := projectmanager.New(utils.CleanProjectType, projectName, destination)
	if err := projectManager.ProcessBaseStructureGeneration(); err != nil {
		return err
	}

	displayEndingMessage(projectName, destination, utils.CleanProjectType)

	return nil
}

func GenerateStructure(projectType utils.ProjectType, projectName, destination string) error {
	if projectType != utils.LiteProjectType && projectType != utils.CleanProjectType {
		return errors.New("invalid project type")
	}

	if !strings.HasSuffix(destination, "/") {
		destination += "/"
	}

	if projectType == utils.CleanProjectType {
		return generateCleanStructure(projectName, destination)
	}

	return generateLiteStructure(projectName, destination)
}

func displayEndingMessage(projectName, destination string, projectType utils.ProjectType) {
	utils.Ui{}.UiSuccess("✅ Génération réussie !")

	utils.Ui{}.UiSection("📂 Project généré", projectName)
	utils.Ui{}.UiSection("📍 Emplacement", destination)
	utils.Ui{}.UiSection("🏗️ Type de projet", projectType)
	fmt.Println()

	utils.Ui{}.UiWarning("🚀 Installez les dépendances")
	utils.Ui{}.UiInfo("   npm install\n")

	utils.Ui{}.UiWarning("🚀 Lancez le projet")
	utils.Ui{}.UiInfo("   npm run dev\n")
}
