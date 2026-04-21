package route

import (
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/utils/typeutils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CleanRouteStrategy struct {
}

func (crs *CleanRouteStrategy) AddRoute(
	controllerNames ControllerNames,
	domainName string,
	completeRouteUrl string,
	method string,
) error {
	controllerDir := GetControllerDirectoryPathByType(filepath.Join(domainName, controllerNames.CleanName), typeutils.CleanProjectType)
	controllerFilePath := filepath.Join(controllerDir, controllerNames.FileName)

	prefix, routeUrl, err := SeparateRoutePrefixFromUrl(completeRouteUrl)
	if err != nil {
		return err
	}

	routeUrl = TransformRouteUrlIntoCleanNotation(routeUrl)
	if routeUrl == "" {
		routeUrl = "/"
	}

	caser := cases.Title(language.Und, cases.NoLower)

	paramName := ExtractParamName(completeRouteUrl)

	data := struct {
		RouteMethod            string
		RoutePrefix            string
		RouteUrl               string
		ControllerFunctionName string
		ParamName              string
	}{
		RouteMethod:            caser.String(strings.ToLower(method)),
		RoutePrefix:            prefix,
		RouteUrl:               routeUrl,
		ControllerFunctionName: controllerNames.FunctionName,
		ParamName:              paramName,
	}

	return CreateControllerFileFromTmpl(
		controllerFilePath,
		typeutils.CleanProjectType,
		method,
		completeRouteUrl,
		&data,
	)
}
