package middleware

import (
	"path/filepath"

	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func MiddlewareAlreadyExists(middlewares []typeutils.CustomMiddlewareData, name string) bool {
	for _, m := range middlewares {
		if m.Name == name {
			return true
		}
	}
	return false
}

func GetMiddlewareFileName(name string) string {
	return name + ".middleware.ts"
}

func GetMiddlewareTemplateFilePath(projectType typeutils.ProjectType) string {
	return filepath.Join(osutils.GetCliRootPath(), "templates", "addcommand", string(projectType), "middleware", "custom.middleware.ts.tmpl")
}

func GetMiddlewareDirectoryPath(projectType typeutils.ProjectType) string {
	projectPath := osutils.GetCurrentDirPath()
	if projectType == typeutils.LiteProjectType {
		return filepath.Join(projectPath, "src", "middlewares")
	}
	return filepath.Join(projectPath, "src", "adapters", "middlewares")
}
