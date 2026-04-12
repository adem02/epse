package service

import (
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func GetServiceTemplatePath(projectType typeutils.ProjectType) string {
	return filepath.Join(
		osutils.GetCliRootPath(),
		"templates",
		"addcommand",
		string(projectType),
		"service",
		"service.ts.tmpl",
	)
}

func GetLiteServiceFilePath(name string) string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "services", name+".service.ts")
}

func GetCleanServiceFilePath(name string) string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "adapters", "services", Capitalize(name)+".service.ts")
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func CleanServiceName(name string) string {
	name = strings.TrimSpace(name)

	name = strings.TrimSuffix(name, "Service")
	name = strings.TrimSuffix(name, "service")
	return name
}
