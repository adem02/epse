package route

import (
	"fmt"

	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type RouteManager struct {
	ControllerName   string
	CompleteRouteUrl string
	Method           string
	DomainName       string
}

type ControllerNames struct {
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

	projectType := typeutils.ProjectType(configData.ProjectType)
	routeStrategy := GetRouteStrategy(projectType)
	controllerNames := GenerateControllerNamesByType(rm.ControllerName, projectType)

	if err := routeStrategy.AddRoute(
		controllerNames,
		rm.DomainName,
		rm.CompleteRouteUrl,
		rm.Method,
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
