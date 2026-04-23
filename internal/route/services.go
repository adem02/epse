package route

import (
	"fmt"
	"strings"

	"github.com/adem02/epse/internal/common"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func GetControllerDirectoryPathByType(domainName string, projectType typeutils.ProjectType) string {
	if projectType == typeutils.LiteProjectType {
		return common.GetFileOrDirectoryPathFromSrcPath("controllers", domainName)
	}

	return common.GetFileOrDirectoryPathFromSrcPath("adapters", "controllers", domainName)
}

func GetLiteRouteIndexFilePath() string {
	return common.GetFileOrDirectoryPathFromSrcPath("routes", "index.ts")
}

func GetLiteNewRouteFilePath(domainName string) string {
	return common.GetFileOrDirectoryPathFromSrcPath("routes", domainName+".routes.ts")
}

func SeparateRoutePrefixFromUrl(completedRouteUrl string) (prefix string, routeUrl string, err error) {
	trimmedUrl := strings.Trim(completedRouteUrl, "/")
	splittedRouteUrl := strings.Split(trimmedUrl, "/")

	if len(splittedRouteUrl) == 0 {
		return "", "", fmt.Errorf("❌ invalid URL: %s", completedRouteUrl)
	}

	prefix = splittedRouteUrl[0]

	if strings.HasPrefix(prefix, ":") {
		return "", "", fmt.Errorf("❌ route prefix cannot be a dynamic variable: %s", prefix)
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

func GenerateControllerNamesByType(controllerName string, projectType typeutils.ProjectType) ControllerNames {
	words := normalizeControllerWords(controllerName)

	pascalName := common.ToPascalCase(words)
	kebabName := common.ToKebabCase(words)

	fileBase := pascalName + ".controller"
	if projectType == typeutils.LiteProjectType {
		fileBase = kebabName + ".controller"
	}

	return ControllerNames{
		CleanName:          pascalName,
		FileName:           fileBase + ".ts",
		FunctionName:       pascalName + "Controller",
		FileNameImportPath: fileBase,
	}
}

func normalizeControllerWords(input string) []string {
	s := strings.TrimSpace(input)
	if s == "" {
		return nil
	}

	lower := strings.ToLower(s)

	if strings.HasSuffix(lower, ".controller.ts") {
		s = s[:len(s)-len(".controller.ts")]
	} else if strings.HasSuffix(lower, ".ts") {
		s = s[:len(s)-len(".ts")]
	}

	lower = strings.ToLower(s)
	if strings.HasSuffix(lower, "-controller") {
		s = s[:len(s)-len("-controller")]
	} else if strings.HasSuffix(lower, "_controller") {
		s = s[:len(s)-len("_controller")]
	} else if strings.HasSuffix(lower, "controller") {
		s = s[:len(s)-len("controller")]
	}

	s = strings.TrimSpace(s)
	s = common.SplitCamelOrPascal(s)

	replacer := strings.NewReplacer("-", " ", "_", " ", ".", " ", "/", " ")
	s = replacer.Replace(s)

	return strings.Fields(strings.ToLower(s))
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
