/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/repository"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
	"github.com/adem02/epse/internal/utils/ui"
	"github.com/spf13/cobra"
)

var repositoryCmd = &cobra.Command{
	Use:   "repository <name>",
	Short: "Generate a repository",
	Long: `Generate a repository for your Express/TypeScript project.

Supported strategies:
  Lite  - Generates a repository in src/repositories/
  Clean - Generates a repository interface in src/useCases/gateway/
          and implementation in src/adapters/gateway/

Usage:
  epse add repository <name>`,
	Run: func(cmd *cobra.Command, args []string) {
		projectPath := osutils.GetCurrentDirPath()
		if !config.ConfigFileExists(projectPath) {
			logutils.Logger{}.Error(fmt.Errorf("❌ fichier de configuration non trouvé"))
			return
		}

		if len(args) < 1 {
			handleAddRepositoryInteractively()
			return
		}

		handleAddRepositoryWithArguments(args)
	},
}

func init() {
	AddCmd.AddCommand(repositoryCmd)
}

func handleAddRepositoryInteractively() {
	var name string

	ui.GetInput(&survey.Input{
		Message: "🗄 Enter the **Repository Name** (e.g., `user`, `product`, `order`):",
	}, &name, func(val interface{}) error {
		str := strings.TrimSpace(val.(string))
		if str == "" {
			return fmt.Errorf("❌ Repository name cannot be empty")
		}
		return nil
	})

	name = strings.TrimSpace(name)

	fmt.Println("\n✅ Configuration Summary:")
	fmt.Println("🔹 Repository Name:", name)

	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to proceed?",
		Default: true,
	}, &confirm)

	if !confirm {
		fmt.Println("❌ Operation canceled.")
		return
	}

	if err := runAddRepository(name); err != nil {
		logutils.Logger{}.Error(err)
	}
}

func handleAddRepositoryWithArguments(args []string) {
	name := strings.TrimSpace(args[0])

	if len(name) < 2 {
		logutils.Logger{}.Error(fmt.Errorf("repository name must have at least 2 characters"))
		return
	}

	if err := runAddRepository(name); err != nil {
		logutils.Logger{}.Error(err)
	}
}

func runAddRepository(name string) error {
	configData, err := config.ReadConfigFileData()
	if err != nil {
		return err
	}

	projectType := typeutils.ProjectType(configData.ProjectType)
	repositoryManager := repository.NewRepositoryManager(name, projectType)

	return repositoryManager.AddRepository()
}
