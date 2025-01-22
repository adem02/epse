/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/adem02/epse/internal/generator"
	"github.com/spf13/cobra"
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
	var projectName, projectType, destination string

	survey.AskOne(&survey.Input{
		Message: "Nom du projet :",
		Default: "api",
	}, &projectName)

	survey.AskOne(&survey.Select{
		Message: "Type de structure :",
		Options: []string{
			"Lite - Node + Express/TypeScript",
			"Clean - Node + Express/TypeScript + TSOA + Clean Architecture",
		},
	}, &projectType)

	survey.AskOne(&survey.Input{
		Message: "Emplacement du projet (défaut : ./) :",
		Default: "./",
	}, &destination)

	if projectType == "Lite - Node + Express/TypeScript" {
		projectType = "lite"
	} else if projectType == "Clean - Node + Express/TypeScript + TSOA + Clean Architecture" {
		projectType = "clean"
	}

	if !strings.HasSuffix(destination, "/") {
		destination += "/"
	}

	err := createProjectStructureByType(projectName, projectType, destination)

	if err != nil {
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

	if !strings.HasSuffix(destination, "/") {
		destination += "/"
	}

	projectType := "lite"

	if clean {
		projectType = "clean"
	}

	err := createProjectStructureByType(projectName, projectType, destination)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func createProjectStructureByType(projectName, projectType, destination string) error {
	if projectType == "clean" || projectType == "lite" {
		newGenerator, err := generator.New(projectName, projectType, destination)

		if err != nil {
			return err
		}

		return newGenerator.GenerateProjectStructure()
	} else {
		return errors.New("❌ invalid project type")
	}
}
