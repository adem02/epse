package auth

import (
	"path/filepath"

	"github.com/adem02/epse/internal/common"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func GetAuthTemplatePath(projectType typeutils.ProjectType, templateName string) string {
	return filepath.Join(
		"addcommand",
		string(projectType),
		"auth",
		templateName,
	)
}

func GetEnvFilePath() string {
	return common.GetFileOrDirectoryFromProjectRootPath(".env")
}

func GetCleanAuthControllerPath(controller, fileName string) string {
	return common.GetFileOrDirectoryPathFromSrcPath("adapters", "controllers", "auth", controller, fileName)
}

func GetCleanAuthUseCasePath(useCase, fileName string) string {
	return common.GetFileOrDirectoryPathFromSrcPath("useCases", "auth", useCase, fileName)
}
