package auth

import (
	"path/filepath"

	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func GetAuthTemplatePath(projectType typeutils.ProjectType, templateName string) string {
	return filepath.Join(
		osutils.GetCliRootPath(),
		"templates",
		"addcommand",
		string(projectType),
		"auth",
		templateName,
	)
}

func GetEnvFilePath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), ".env")
}

func GetLiteJwtConfigPath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "config", "jwt.config.ts")
}

func GetLiteAuthTypesPath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "types", "auth.types.ts")
}

func GetLiteJwtServicePath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "services", "jwt.service.ts")
}

func GetLiteAuthServicePath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "services", "auth.service.ts")
}

func GetLiteAuthMiddlewarePath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "middlewares", "auth.middleware.ts")
}

func GetLiteLoginControllerPath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "controllers", "auth", "Login.controller.ts")
}

func GetLiteRegisterControllerPath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "controllers", "auth", "Register.controller.ts")
}

func GetLiteAuthRoutesPath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "routes", "auth.routes.ts")
}

func GetLiteExpressDtsPath() string {
	return filepath.Join(osutils.GetCurrentDirPath(), "src", "types", "express.d.ts")
}

func GetCleanSrcPath(folders ...string) string {
	parts := append([]string{osutils.GetCurrentDirPath(), "src"}, folders...)
	return filepath.Join(parts...)
}

func GetCleanAuthControllerPath(controller, fileName string) string {
	return GetCleanSrcPath("adapters", "controllers", "auth", controller, fileName)
}

func GetCleanAuthUseCasePath(useCase, fileName string) string {
	return GetCleanSrcPath("useCases", "auth", useCase, fileName)
}
