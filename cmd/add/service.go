/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/service"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
	"github.com/adem02/epse/internal/utils/ui"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service <name>",
	Short: "Generate a service",
	Long: `Generate a service for your Express/TypeScript project.

Supported strategies:
  Lite  - Generates a service in src/services/
  Clean - Generates an injectable service in src/adapters/services/

Usage:
  epse add service <name>`,
	Run: func(cmd *cobra.Command, args []string) {
		if !config.ConfigFileExists() {
			logutils.Logger{}.Error(fmt.Errorf("❌ fichier de configuration non trouvé"))
			return
		}

		if len(args) < 1 {
			handleAddServiceInteractively()
			return
		}

		handleAddServiceWithArguments(args)
	},
}

func init() {
	AddCmd.AddCommand(serviceCmd)
}

func handleAddServiceInteractively() {
	var name string

	ui.GetInput(&survey.Input{
		Message: "🔧 Enter the **Service Name** (e.g., `user`, `email`, `payment`):",
	}, &name, func(val interface{}) error {
		str := strings.TrimSpace(val.(string))
		if str == "" {
			return fmt.Errorf("❌ Service name cannot be empty")
		}
		return nil
	})

	name = strings.TrimSpace(name)

	if len(name) < 2 {
		logutils.Logger{}.Error(fmt.Errorf("service name must have at least 2 characters"))
		return
	}

	fmt.Println("\n✅ Configuration Summary:")
	fmt.Println("🔹 Service Name:", name)

	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to proceed?",
		Default: true,
	}, &confirm)

	if !confirm {
		fmt.Println("❌ Operation canceled.")
		return
	}

	if err := runAddService(name); err != nil {
		logutils.Logger{}.Error(err)
	}
}

func handleAddServiceWithArguments(args []string) {
	name := strings.TrimSpace(args[0])

	if len(name) < 2 {
		logutils.Logger{}.Error(fmt.Errorf("service name must have at least 2 characters"))
		return
	}

	fmt.Println("\n✅ Configuration Summary:")
	fmt.Println("🔹 Service Name:", name)

	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to proceed?",
		Default: true,
	}, &confirm)

	if !confirm {
		fmt.Println("❌ Operation canceled.")
		return
	}

	if err := runAddService(name); err != nil {
		logutils.Logger{}.Error(err)
	}
}

func runAddService(name string) error {
	configData, err := config.ReadConfigFileData()
	if err != nil {
		return err
	}

	projectType := typeutils.ProjectType(configData.ProjectType)
	names := service.GenerateServiceNamesByType(name, projectType)
	serviceManager := service.NewServiceManager(names, projectType)

	return serviceManager.AddService()
}
