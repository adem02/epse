/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/adem02/epse/cmd/add"
	"github.com/adem02/epse/cmd/generate"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "epse",
	Short: "A brief description of your application",
	Long: `EPSE: A CLI to generate Node.js, Express, and TypeScript project structures.

Usage:
  epse [command]

Available Commands:
  generate    Generate a new project structure
  help        Display help for available commands`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to EPSE! Use 'epse help' to see available options.")
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(generate.GenerateCmd)
	RootCmd.AddCommand(add.AddCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.epse.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
