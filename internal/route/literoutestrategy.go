package route

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/utils/typeutils"
)

type LiteRouteStrategy struct {
}

func (lrs *LiteRouteStrategy) AddRoute(
	controllerNames ControllerNames,
	domainName string,
	completeRouteUrl string,
	method string,
	authMiddleware,
	adminAuthMiddleware bool,
) error {
	controllerDir := GetControllerDirectoryPathByType(domainName, typeutils.LiteProjectType)
	controllerFilePath := filepath.Join(controllerDir, controllerNames.FileName)

	data := struct {
		ControllerFunc string
	}{
		ControllerFunc: controllerNames.FunctionName,
	}

	if err := CreateControllerFileFromTmpl(
		controllerFilePath,
		typeutils.LiteProjectType,
		method,
		completeRouteUrl,
		&data,
	); err != nil {
		return err
	}

	prefix, routeUrl, err := SeparateRoutePrefixFromUrl(completeRouteUrl)
	if err != nil {
		return err
	}

	if routeUrl == "" {
		routeUrl = "/"
	}

	routeFilePath := GetLiteNewRouteFilePath(domainName)
	bytes, err := os.ReadFile(routeFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			bytes = []byte(NewRouteFileStringTmpl(
				controllerNames.FunctionName,
				controllerNames.FileName,
				domainName,
				routeUrl,
				method,
			))

			if err := os.WriteFile(routeFilePath, bytes, 0644); err != nil {
				return err
			}

			return updateIndexRouteFile(domainName, prefix)
		} else {
			return err
		}
	}

	content := string(bytes)

	newRoute := fmt.Sprintf(`router.%s('%s', %s);`, strings.ToLower(method), routeUrl, controllerNames.FunctionName)
	importStatement := fmt.Sprintf(`import { %s } from '@/controllers/%s/%s';`, controllerNames.FunctionName, domainName, controllerNames.FileNameImportPath)

	AddImportStatementInFile(&content, importStatement)
	if !strings.Contains(content, newRoute) {
		InsertContentBeforeStatementInFile(&content, "export default router;", newRoute)
	}

	if err := os.WriteFile(routeFilePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("❌ Error while writing in the route file: %w", err)
	}

	return nil
}

func updateIndexRouteFile(domainName, prefix string) error {
	indexFilePath := GetLiteRouteIndexFilePath(domainName)
	bytes, err := os.ReadFile(indexFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("❌ File `routes/index.ts` not found in project")
		}
		return fmt.Errorf("❌ Error while reading the file `routes/index.ts`: %w", err)
	}

	content := string(bytes)

	routeVarName := fmt.Sprintf("%sRoutes", domainName)
	importStatement := fmt.Sprintf(`import %s from './%s.routes';`, routeVarName, domainName)
	newRouteUse := fmt.Sprintf("router.use('/%s', %s);", prefix, routeVarName)

	AddImportStatementInFile(&content, importStatement)
	if !strings.Contains(content, newRouteUse) {
		InsertContentBeforeStatementInFile(&content, "export default router;", newRouteUse)
	}

	if err := os.WriteFile(indexFilePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("❌ Error while writing in the route file: %w", err)
	}

	return nil
}
