/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/adem02/epse/internal/generator"
	"github.com/adem02/epse/internal/utils"
	"github.com/spf13/cobra"
	"os"
)

var lite, clean bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
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
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().BoolVar(&lite, "lite", false, "Générer un projet minimaliste (Lite)")
	generateCmd.Flags().BoolVar(&clean, "clean", false, "Générer un projet plus conséquent en avec tsoa (Clean)")
}

func runInteractive() {
	var projectName, destination string
	var projectType utils.ProjectType

	getInput(&survey.Input{
		Message: "Nom du projet :",
		Default: "api",
	}, &projectName)

	getInput(&survey.Select{
		Message: "Type de structure :",
		Options: []string{
			"Lite - Node + Express/TypeScript",
			"Clean - Node + Express/TypeScript + TSOA + Clean Architecture",
		},
	}, &projectType)

	getInput(&survey.Input{
		Message: "Emplacement du projet (défaut : ./) :",
		Default: "./",
	}, &destination)

	if projectType == "Lite - Node + Express/TypeScript" {
		projectType = utils.LiteProjectType
	} else if projectType == "Clean - Node + Express/TypeScript + TSOA + Clean Architecture" {
		projectType = utils.CleanProjectType
	}

	if err := createProjectStructureByType(projectType, projectName, destination); err != nil {
		fmt.Println(err)
		return
	}
}

func runWithArguments(args []string) {
	projectName := args[0]
	destination := "./"

	if len(args) > 1 {
		destination = args[1]
	}

	projectType := utils.LiteProjectType

	if clean {
		projectType = utils.CleanProjectType
	}

	if err := createProjectStructureByType(projectType, projectName, destination); err != nil {
		utils.Ui{}.UiError(err)
		return
	}
}

func createProjectStructureByType(projectType utils.ProjectType, projectName, destination string) error {
	return generator.GenerateStructure(projectType, projectName, destination)
}

func getInput(prompt survey.Prompt, response interface{}) {
	if err := survey.AskOne(prompt, response); err != nil {
		utils.Ui{}.UiError(fmt.Errorf("\n  interruption détectée. Fermeture...\n"))
		os.Exit(1)
	}
}
