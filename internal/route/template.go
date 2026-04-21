package route

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/templates"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func NewRouteFileStringTmpl(controllerFunctionName, controllerFileName, domainName, routeUrl, method string) string {
	return fmt.Sprintf(`import { Router } from 'express';
import { %s } from '@/controllers/%s/%s';

const router = Router();

router.%s('%s', %s);

export default router;
`,
		controllerFunctionName, domainName, strings.TrimSuffix(controllerFileName, ".ts"),
		strings.ToLower(method), routeUrl, controllerFunctionName)
}

func CreateControllerFileFromTmpl(
	controllerFilePath string,
	projectType typeutils.ProjectType,
	method string,
	completeRouteUrl string,
	data interface{},
) error {
	if osutils.FileOrDirectoryExists(controllerFilePath) {
		logutils.Logger{}.Warning(fmt.Sprintf("⚠️ Le contrôleur existe déjà : %s", controllerFilePath))
		return nil
	}

	if err := osutils.CreateDirectory(filepath.Dir(controllerFilePath)); err != nil {
		return err
	}

	_, routeUrl, err := SeparateRoutePrefixFromUrl(completeRouteUrl)
	if err != nil {
		return err
	}

	hasParam := RouteHasPathParam(routeUrl)

	templateName := GetControllerTemplateByMethod(method, hasParam)
	templatePath := filepath.Join(
		"addcommand",
		string(projectType),
		"route",
		templateName,
	)

	return osutils.CreateFileFromTmpl(
		templates.FS,
		templatePath,
		controllerFilePath,
		data,
	)
}
