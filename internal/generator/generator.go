package generator

import (
	"errors"
	"fmt"

	"github.com/fatih/color"

	"github.com/adem02/epse/internal/projectmanager"
	"github.com/adem02/epse/internal/utils"
)

type Generator struct {
	ProjectType utils.ProjectType
	ProjectName string
	Destination string
}

var section = color.New(color.Bold, color.FgGreen).SprintFunc()
var info = color.New(color.FgHiBlue).SprintFunc()
var success = color.New(color.FgGreen).SprintFunc()
var warning = color.New(color.FgYellow).SprintFunc()

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

func (g Generator) GenerateProjectStructure() error {
	var err error = nil

	if g.ProjectType == utils.LiteProjectType {
		err = generateLiteStructure(g.ProjectName, g.Destination)
	} else if g.ProjectType == utils.CleanProjectType {
		err = generateCleanStructure(g.ProjectName, g.Destination)
	} else {
		return errors.New("failed to generate project structure of unknown type")
	}

	return err
}

func New(projectType utils.ProjectType, projectName, destination string) (Generator, error) {
	if projectType != utils.LiteProjectType && projectType != utils.CleanProjectType {
		return Generator{}, errors.New("invalid project type")
	}

	return Generator{
		ProjectType: projectType,
		ProjectName: projectName,
		Destination: destination,
	}, nil
}

func displayEndingMessage(projectName, destination string, projectType utils.ProjectType) {
	fmt.Println(success("✅ Génération réussie !"))
	fmt.Printf("   📂 %s: %s\n", section("Projet généré"), info(projectName))
	fmt.Printf("   📍 %s: %s\n", section("Emplacement"), info(destination))
	fmt.Printf("   🏗️ %s: %s\n\n", section("Type de projet"), info(projectType))

	fmt.Println(warning("🚀 Installez les dépendances"))
	fmt.Println(info("   npm install\n"))

	fmt.Println(warning("🚀 Lancez le projet"))
	fmt.Println(info("   npm run dev\n"))
}
