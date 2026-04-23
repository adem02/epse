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
	Short: "Generate a new project structure",
	Long: `Generate a new project structure based on the project name and destination.
If destination is not specified, the project will be generated in the current directory.
The generated project is based on either a minimal template or a more complete one depending on selected options.`,
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
	GenerateCmd.Flags().BoolVar(&lite, "lite", false, "Generate a minimal project (Lite)")
	GenerateCmd.Flags().BoolVar(&clean, "clean", false, "Generate a complete project with TSOA (Clean)")
}

func runInteractive() {
	var projectName, destination string
	var projectType typeutils.ProjectType

	ui.GetInput(&survey.Input{
		Message: "Project name:",
		Default: "api",
	}, &projectName, survey.Required)

	ui.GetInput(&survey.Select{
		Message: "Project structure type:",
		Options: []string{
			"Lite - Node + Express/TypeScript",
			"Clean - Node + Express/TypeScript + TSOA + Clean Architecture",
		},
	}, &projectType, survey.Required)

	ui.GetInput(&survey.Input{
		Message: "Project location (default: ./):",
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
