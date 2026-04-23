/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/adem02/epse/internal/auth"
	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Generate a complete JWT authentication system",
	Long: `Generate a complete JWT authentication system for your Express/TypeScript project.

Supported strategies:
  Lite  - Generates JWT auth with login, register, middleware and routes
  Clean - Generates JWT auth with login, register, use cases, controllers, gateway and TSOA authentication

Usage:
  epse add auth`,
	Run: func(cmd *cobra.Command, args []string) {
		if !config.ConfigFileExists() {
			logutils.Logger{}.Error(fmt.Errorf("❌ configuration file not found"))
			return
		}

		if err := runAddAuth(); err != nil {
			logutils.Logger{}.Error(err)
		}
	},
}

func init() {
	AddCmd.AddCommand(authCmd)
}

func runAddAuth() error {
	configData, err := config.ReadConfigFileData()
	if err != nil {
		return err
	}

	projectType := typeutils.ProjectType(configData.ProjectType)

	fmt.Println("\n✅ Configuration Summary:")
	fmt.Println("🔹 Feature: JWT Authentication")
	fmt.Println("🔹 Project Type:", projectType)

	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to proceed with this configuration?",
		Default: true,
	}, &confirm)

	if !confirm {
		fmt.Println("❌ Operation canceled.")
		return nil
	}

	authManager := auth.NewAuthManager(projectType)

	return authManager.AddAuth()
}
