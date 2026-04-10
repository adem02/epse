package add

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/middleware"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
	"github.com/adem02/epse/internal/utils/ui"
	"github.com/spf13/cobra"
)

var MiddlewareCmd = &cobra.Command{
	Use:   "middleware <name>",
	Short: "Generate a custom middleware",
	Long:  `Generate a custom Express middleware file in src/middlewares/`,
	Run: func(cmd *cobra.Command, args []string) {
		projectPath := osutils.GetCurrentDirPath()
		if !config.ConfigFileExists(projectPath) {
			logutils.Logger{}.Error(fmt.Errorf("❌ fichier de configuration non trouvé"))
			return
		}

		if len(args) < 1 {
			handleAddMiddlewareInteractively()
			return
		}

		handleAddMiddlewareWithArguments(args)
	},
}

func init() {
	AddCmd.AddCommand(MiddlewareCmd)
}

func handleAddMiddlewareInteractively() {
	var name string

	ui.GetInput(&survey.Input{
		Message: "🛡 Enter the **Middleware Name** (e.g., `auth`, `logger`, `rate-limit`):",
	}, &name, func(val interface{}) error {
		str := strings.TrimSpace(val.(string))
		if str == "" {
			return fmt.Errorf("❌ Middleware name cannot be empty")
		}
		return nil
	})

	name = strings.TrimSpace(name)

	fmt.Println("\n✅ Configuration Summary:")
	fmt.Println("🔹 Middleware Name:", name)

	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to proceed?",
		Default: true,
	}, &confirm)

	if !confirm {
		fmt.Println("❌ Operation canceled.")
		return
	}

	fmt.Println()
	logutils.Logger{}.Info("🚀 Generating middleware...")

	if err := runAddMiddleware(name); err != nil {
		logutils.Logger{}.Error(err)
	}
}

func handleAddMiddlewareWithArguments(args []string) {
	name := strings.TrimSpace(args[0])

	if len(name) < 2 {
		logutils.Logger{}.Error(fmt.Errorf("middleware name must have at least 2 characters"))
		return
	}

	fmt.Println("\n✅ Configuration Summary:")
	fmt.Println("🔹 Middleware Name:", name)

	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to proceed?",
		Default: true,
	}, &confirm)

	if !confirm {
		fmt.Println("❌ Operation canceled.")
		return
	}

	fmt.Println()
	logutils.Logger{}.Info("🚀 Generating middleware...")

	if err := runAddMiddleware(name); err != nil {
		logutils.Logger{}.Error(err)
	}
}

func runAddMiddleware(name string) error {
	configData, err := config.ReadConfigFileData()
	if err != nil {
		return err
	}

	projectType := typeutils.ProjectType(configData.ProjectType)
	middlewareManager := middleware.NewMiddlewareManager(name, projectType)

	return middlewareManager.AddMiddleware()
}
