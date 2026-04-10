package route

import (
	"fmt"
	"strings"

	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type RouteManager struct {
	ControllerName      string
	CompleteRouteUrl    string
	Method              string
	DomainName          string
	AuthMiddleware      bool
	AdminAuthMiddleware bool
}

type ControllerNames struct {
	OriginalName       string
	CleanName          string
	FileName           string
	FunctionName       string
	FileNameImportPath string
}

func (rm *RouteManager) AddRoute() error {
	configData, err := config.ReadConfigFileData()
	if err != nil {
		return err
	}

	routeStrategy := GetRouteStrategy(typeutils.ProjectType(configData.ProjectType))
	controllerNames := GenerateControllerNames(rm.ControllerName)

	if err := routeStrategy.AddRoute(
		controllerNames,
		rm.DomainName,
		rm.CompleteRouteUrl,
		rm.Method,
		rm.AuthMiddleware,
		rm.AdminAuthMiddleware,
	); err != nil {
		return err
	}

	prefix, _, err := SeparateRoutePrefixFromUrl(rm.CompleteRouteUrl)
	if err != nil {
		return err
	}

	if !config.DomainNameAlreadyExistsInRoutes(configData.Routes, rm.DomainName) {
		if err := config.AddNewRouteInConfigFile(rm.DomainName, prefix, configData); err != nil {
			return err
		}
	}

	rm.displaySuccess(controllerNames, configData.ProjectType)

	return nil
}

func NewRouteManager(
	controllerName,
	completedRouteUrl,
	method,
	domainName string,
) *RouteManager {
	return &RouteManager{
		ControllerName:   controllerName,
		CompleteRouteUrl: completedRouteUrl,
		Method:           method,
		DomainName:       domainName,
	}
}

func cleanControllerName(name string) string {
	lowerredName := strings.ToLower(strings.TrimSpace(name))

	if strings.HasSuffix(lowerredName, ".controller.ts") {
		lowerredName = lowerredName[:len(name)-len(".controller.ts")]
	} else if strings.HasSuffix(name, ".ts") {
		lowerredName = lowerredName[:len(name)-len(".ts")]
	}

	lowerredName = strings.TrimSuffix(lowerredName, "controller")
	name = name[:len(lowerredName)]

	caser := cases.Title(language.Und, cases.NoLower)
	return caser.String(name)
}

func (rm *RouteManager) displaySuccess(controllerNames ControllerNames, projectType string) {
	fmt.Println()
	logutils.Logger{}.Success("✅ Route added successfully!")
	fmt.Println()
	logutils.Logger{}.Info("📁 Files:")

	controllerPath := rm.getControllerPath(controllerNames, projectType)
	fmt.Printf("  ✓ %s\n", controllerPath)

	if projectType == "lite" {
		routePath := rm.getRoutePath(projectType)
		fmt.Printf("  ✓ %s (updated)\n", routePath)
		// TODO: - Afficher "epseconfig.json (updated)" seulement si domain nouveau
		// TODO: - Afficher "epseconfig.json (updated)" seulement si domain nouveau
		fmt.Printf("  ✓ src/routes/index.ts (updated)\n")
	} else {
		fmt.Printf("  ✓ TSOA routes (auto-generated on build)\n")
	}

	fmt.Printf("  ✓ epseconfig.json (updated)\n")
	fmt.Println()
}

func (rm *RouteManager) getControllerPath(controllerNames ControllerNames, projectType string) string {
	if projectType == "lite" {
		return fmt.Sprintf("src/controllers/%s/%s", rm.DomainName, controllerNames.FileName)
	}
	return fmt.Sprintf("src/adapters/controllers/%s/%s/%s",
		rm.DomainName, controllerNames.CleanName, controllerNames.FileName)
}

func (rm *RouteManager) getRoutePath(projectType string) string {
	if projectType == "lite" {
		return fmt.Sprintf("src/routes/%s.routes.ts", rm.DomainName)
	}
	return ""
}
