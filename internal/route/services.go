package route

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func GetControllerTemplateFilePath(projectType typeutils.ProjectType) string {
	return filepath.Join(osutils.GetCliRootPath(), "templates", "addcommand", string(projectType), "route", "generatedcontroller.ts.tmpl")
}

func GetControllerDirectoryPathByType(domainName string, projectType typeutils.ProjectType) string {
	projectPath := osutils.GetCurrentDirPath()
	if projectType == typeutils.LiteProjectType {
		return filepath.Join(projectPath, "src", "controllers", domainName)
	}

	return filepath.Join(projectPath, "src", "adapters", "controllers", domainName)
}

func GetLiteRouteIndexFilePath(domainName string) string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "routes", "index.ts")
}

func GetLiteNewRouteFilePath(domainName string) string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "routes", domainName+".routes.ts")
}

func SeparateRoutePrefixFromUrl(completedRouteUrl string) (prefix string, routeUrl string, err error) {
	trimmedUrl := strings.Trim(completedRouteUrl, "/")
	splittedRouteUrl := strings.Split(trimmedUrl, "/")

	if len(splittedRouteUrl) == 0 {
		return "", "", fmt.Errorf("❌ URL invalide : %s", completedRouteUrl)
	}

	prefix = splittedRouteUrl[0]

	if strings.HasPrefix(prefix, ":") {
		return "", "", fmt.Errorf("❌ Un préfixe de route ne peut pas être une variable dynamique : %s", prefix)
	}

	routeUrl = strings.Join(splittedRouteUrl[1:], "/")

	if routeUrl == "" {
		routeUrl = "/"
	}

	return prefix, routeUrl, nil
}

func TransformRouteUrlIntoCleanNotation(routeUrl string) string {
	trimmedUrl := strings.Trim(routeUrl, "/")
	splittedRouteUrl := strings.Split(trimmedUrl, "/")

	for index, param := range splittedRouteUrl {
		if strings.Contains(param, ":") {
			splittedRouteUrl[index] = strings.Replace(param, param, "{"+param[1:]+"}", 1)
		}
	}

	return strings.Join(splittedRouteUrl, "/")
}

func GenerateControllerNames(controllerName string) ControllerNames {
	originalName := controllerName
	controllerName = cleanControllerName(controllerName)

	fileName := controllerName + ".controller.ts"
	functionName := controllerName + "Controller"
	importPath := strings.TrimSuffix(fileName, ".ts")

	return ControllerNames{
		OriginalName:       originalName,
		CleanName:          controllerName,
		FileName:           fileName,
		FunctionName:       functionName,
		FileNameImportPath: importPath,
	}
}

func RouteHasPathParam(routeUrl string) bool {
	return strings.Contains(routeUrl, ":")
}

func GetControllerTemplateByMethod(method string, hasPathParam bool) string {
	method = strings.ToUpper(method)

	switch method {
	case "GET":
		if hasPathParam {
			return "get-with-param.controller.ts.tmpl"
		}
		return "get.controller.ts.tmpl"
	case "POST":
		return "post.controller.ts.tmpl"
	case "PUT":
		return "put.controller.ts.tmpl"
	case "PATCH":
		return "patch.controller.ts.tmpl"
	case "DELETE":
		return "delete.controller.ts.tmpl"
	default:
		return "generic.controller.ts.tmpl"
	}
}

func ExtractParamName(routeUrl string) string {
	parts := strings.Split(routeUrl, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			return strings.TrimPrefix(part, ":")
		}
	}
	return "id"
}
