/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/route"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/ui"
	"github.com/spf13/cobra"
)

var authenticated, admin, crud bool
var actionMethod, controllerName string
var actionMethods = []string{
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
}

// routeCmd represents the route command
var RouteCmd = &cobra.Command{
	Use:   "route <domaineName> <routeUrl>",
	Short: "Will generate required files for a give route path",
	Long:  `API route, you must provide details the module name, the path, controller name`,
	Run: func(cmd *cobra.Command, args []string) {
		projectPath := osutils.GetCurrentDirPath()
		if !config.ConfigFileExists(projectPath) {
			logutils.Logger{}.Error(fmt.Errorf("❌ fichier de configuration non trouvé"))
			return
		}

		if len(args) < 1 {
			handleAddRouteInteractively(projectPath)
			return
		}

		handleAddRouteWithArguments(args)

	},
}

func init() {
	AddCmd.AddCommand(RouteCmd)

	RouteCmd.Flags().StringVar(&actionMethod, "method", "GET", "Méthode HTTP (GET, POST, PUT...)")
	RouteCmd.Flags().StringVar(&controllerName, "controller", "", "Nom du controller")
	//RouteCmd.MarkFlagRequired("controller")
	RouteCmd.Flags().BoolVar(&authenticated, "authenticated", false, "Authorized route for authenticated user only")
	RouteCmd.Flags().BoolVar(&admin, "admin", false, "Authorized route for admin user only")
	RouteCmd.Flags().BoolVar(&crud, "crud", false, "Generate complete CRUD routes (GET, POST, PUT, DELETE)")
}

func handleAddRouteInteractively(configFile string) {
	var domainName, routePath, controllerName, actionMethod string

	ui.GetInput(&survey.Input{
		Message: "🌍 Enter the **Domain Name** (e.g., `user`, `product`, `order`):",
	}, &domainName, survey.Required)

	ui.GetInput(&survey.Input{
		Message: "📍 Enter the **Route URI** (e.g., `/users`, `/products/:id`):",
	}, &routePath, func(val interface{}) error {
		str := strings.TrimSpace(val.(string))
		if !strings.HasPrefix(str, "/") {
			return fmt.Errorf("❌ The route must start with `/` (e.g., `/users`)")
		}
		return nil
	})

	var useCrud bool
	survey.AskOne(&survey.Confirm{
		Message: "Generate complete CRUD routes? (GET all, GET by id, POST, PUT, DELETE)",
		Default: false,
	}, &useCrud)

	if useCrud {
		fmt.Println("\n✅ CRUD Configuration:")
		fmt.Println("🔹 Domain Name:", domainName)
		fmt.Println("🔹 Base Path:", routePath)
		fmt.Println("🔹 Will generate 5 routes")

		var confirm bool
		survey.AskOne(&survey.Confirm{
			Message: "Do you want to proceed?",
			Default: true,
		}, &confirm)

		if confirm {
			generateCRUDRoutes(domainName, routePath)
		} else {
			fmt.Println("❌ Operation canceled.")
		}
		return
	}

	ui.GetInput(&survey.Select{
		Message: "⚡ Select the **HTTP Method**:",
		Options: actionMethods,
	}, &actionMethod, survey.Required)

	ui.GetInput(&survey.Input{
		Message: "🎛 Enter the **Controller Name** (e.g., `GetUserController`, `GetUser`):",
	}, &controllerName, func(val interface{}) error {
		str := strings.TrimSpace(val.(string))
		if str == "" {
			return fmt.Errorf("❌ Controller name cannot be empty")
		}
		return nil
	})

	fmt.Println("\n✅ Configuration Summary:")
	fmt.Println("🔹 Domain Name:", domainName)
	fmt.Println("🔹 Route Path:", routePath)
	fmt.Println("🔹 HTTP Method:", actionMethod)
	fmt.Println("🔹 Controller Name:", controllerName)

	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to proceed with this configuration?",
		Default: true,
	}, &confirm)

	if confirm {
		fmt.Println()
		logutils.Logger{}.Info("🚀 Proceeding to generate the route and controller...")

		routeManager := route.NewRouteManager(
			controllerName,
			routePath,
			actionMethod,
			domainName,
		)
		if err := routeManager.AddRoute(); err != nil {
			logutils.Logger{}.Error(err)
		}

	} else {
		fmt.Println("❌ Operation canceled.")
	}
}

func handleAddRouteWithArguments(args []string) {
	domainName := args[0]
	routePath := args[1]

	if crud {
		generateCRUDRoutes(domainName, routePath)
		return
	}

	if len(controllerName) < 4 {
		logutils.Logger{}.Error(fmt.Errorf("controller name must have at least 4 characters"))
		return
	}

	if !slices.Contains(actionMethods, actionMethod) {
		logutils.Logger{}.Error(fmt.Errorf("action %s not allowed", actionMethod))
		return
	}

	fmt.Println("\n✅ Configuration Summary:")
	fmt.Println("🔹 Domain Name:", domainName)
	fmt.Println("🔹 Route Path:", routePath)
	fmt.Println("🔹 HTTP Method:", actionMethod)
	fmt.Println("🔹 Controller Name:", controllerName)

	var confirm bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to proceed with this configuration?",
		Default: true,
	}, &confirm)

	if confirm {
		fmt.Println()
		logutils.Logger{}.Info("🚀 Proceeding to generate the route and controller...")

		routeManager := route.NewRouteManager(
			controllerName,
			routePath,
			actionMethod,
			domainName,
		)
		if err := routeManager.AddRoute(); err != nil {
			logutils.Logger{}.Error(err)
		}

	} else {
		fmt.Println("❌ Operation canceled.")
	}
}

func generateCRUDRoutes(domainName, routeBasePath string) {
	capitalizedDomain := capitalize(domainName)

	crudRoutes := []struct {
		method         string
		path           string
		controllerName string
	}{
		{"GET", routeBasePath, "GetAll" + capitalizedDomain + "s"},
		{"GET", routeBasePath + "/:id", "Get" + capitalizedDomain + "ById"},
		{"POST", routeBasePath, "Create" + capitalizedDomain},
		{"PUT", routeBasePath + "/:id", "Update" + capitalizedDomain + "ById"},
		{"DELETE", routeBasePath + "/:id", "Delete" + capitalizedDomain + "ById"},
	}

	fmt.Println()
	logutils.Logger{}.Info("🔄 Generating CRUD routes...")
	fmt.Println()

	for _, r := range crudRoutes {
		routeManager := route.NewRouteManager(
			r.controllerName,
			r.path,
			r.method,
			domainName,
		)

		if err := routeManager.AddRoute(); err != nil {
			logutils.Logger{}.Error(err)
			return
		}
	}

	fmt.Println()
	logutils.Logger{}.Success("✅ All CRUD routes generated successfully!")
	fmt.Println()
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
