/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new resource (route, middleware, etc.)",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
