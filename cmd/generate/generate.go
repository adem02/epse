/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package generate

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/adem02/epse/internal/project"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
	"github.com/adem02/epse/internal/utils/ui"
	"github.com/spf13/cobra"
)

var lite, clean bool

// GenerateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate <project-name> [destination]",
	Short: "Génère une nouvelle structure de projet",
	Long: `Génère une nouvelle structure de projet en fonction du nom du projet et de la destination.
Si la destination n'est pas spécifiée, le projet sera généré dans le répertoire courant.
Le projet généré sera basé sur un template minimaliste ou plus conséquent en fonction des options.`,
	// Args: cobra.(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			runInteractive()
			return
		}

		runWithArguments(args)
	},
}

func init() {
	GenerateCmd.Flags().BoolVar(&lite, "lite", false, "Générer un projet minimaliste (Lite)")
	GenerateCmd.Flags().BoolVar(&clean, "clean", false, "Générer un projet plus conséquent en avec tsoa (Clean)")
}

func runInteractive() {
	var projectName, destination string
	var projectType typeutils.ProjectType

	ui.GetInput(&survey.Input{
		Message: "Nom du projet :",
		Default: "api",
	}, &projectName, survey.Required)

	ui.GetInput(&survey.Select{
		Message: "Type de structure :",
		Options: []string{
			"Lite - Node + Express/TypeScript",
			"Clean - Node + Express/TypeScript + TSOA + Clean Architecture",
		},
	}, &projectType, survey.Required)

	ui.GetInput(&survey.Input{
		Message: "Emplacement du projet (défaut : ./) :",
		Default: "./",
	}, &destination, survey.Required)

	if projectType == "Lite - Node + Express/TypeScript" {
		projectType = typeutils.LiteProjectType
	} else if projectType == "Clean - Node + Express/TypeScript + TSOA + Clean Architecture" {
		projectType = typeutils.CleanProjectType
	}

	if err := createProjectStructureByType(projectType, projectName, destination); err != nil {
		logutils.Logger{}.Error(err)
		return
	}
}

func runWithArguments(args []string) {
	projectName := args[0]
	destination := "./"

	if len(args) > 1 {
		destination = args[1]
	}

	projectType := typeutils.LiteProjectType

	if clean {
		projectType = typeutils.CleanProjectType
	}

	if err := createProjectStructureByType(projectType, projectName, destination); err != nil {
		logutils.Logger{}.Error(err)
		return
	}
}

func createProjectStructureByType(projectType typeutils.ProjectType, projectName, destination string) error {
	projectManager := project.New(projectType, projectName, destination)
	return projectManager.Generate()
}
