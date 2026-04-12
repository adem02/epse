package repository

import (
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func GetRepositoryTemplatePath(projectType typeutils.ProjectType, templateName string) string {
	return filepath.Join(
		osutils.GetCliRootPath(),
		"templates",
		"addcommand",
		string(projectType),
		"repository",
		templateName,
	)
}

func GetLiteRepositoryFilePath(name string) string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "repositories", strings.ToLower(name)+".repository.ts")
}

func GetCleanRepositoryFilePath(name string) string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "adapters", "gateway", Capitalize(name)+".repository.ts")
}

func GetCleanRepositoryInterfaceFilePath(name string) string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "useCases", "gateway", Capitalize(name)+".repository.interface.ts")
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func CleanRepositoryName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.TrimSuffix(name, "Repository")
	name = strings.TrimSuffix(name, "repository")
	return name
}
