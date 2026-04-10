package project

import (
	"fmt"
	"path/filepath"

	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

var ProjectStructureMappedByProjectType = map[typeutils.ProjectType][]string{
	typeutils.LiteProjectType: {
		"test", "src/controllers/health", "src/models", "src/middlewares",
		"src/repositories", "src/types", "src/services", "src/utils",
		"src/routes", "src/config", "src/errors",
	},
	typeutils.CleanProjectType: {
		"test", "src/useCases", "src/utilities",
		"src/adapters/controllers/health", "src/adapters/gateway",
		"src/adapters/middlewares", "src/adapters/services",
		"src/entities/error", "src/entities/types", "src/entities/logger",
		"src/frameworks/tsoa/services",
	},
}

func GetProjectStructureByType(projectType typeutils.ProjectType) ([]string, error) {
	if projectType != typeutils.LiteProjectType && projectType != typeutils.CleanProjectType {
		return nil, fmt.Errorf("invalid project type")
	}

	return ProjectStructureMappedByProjectType[projectType], nil
}

func GetLiteFilesTemplatesPaths() map[string]string {
	liteTemplatesPath := typeutils.GetLiteTemplatesPath + "/"

	routesPath,
		configPath,
		middlewaresPath,
		controllersPath,
		errorsPath := typeutils.SrcPath+"routes/",
		typeutils.SrcPath+"config/",
		typeutils.SrcPath+"middlewares/",
		typeutils.SrcPath+"controllers/",
		typeutils.SrcPath+"errors/"

	srcTmplPath, routesTmplPath, controllersTmplPath,
		configTmplPath, middlewaresTmplPath, errorsTmplPath :=
		liteTemplatesPath+typeutils.SrcPath, liteTemplatesPath+routesPath, liteTemplatesPath+controllersPath,
		liteTemplatesPath+configPath, liteTemplatesPath+middlewaresPath,
		liteTemplatesPath+errorsPath

	return map[string]string{
		"package.json":      liteTemplatesPath + "package.json.tmpl",
		"README.md":         liteTemplatesPath + "README.md.tmpl",
		"tsconfig.json":     liteTemplatesPath + "tsconfig.json.tmpl",
		".env":              liteTemplatesPath + ".env.tmpl",
		".prettierrc":       liteTemplatesPath + ".prettierrc.tmpl",
		".prettierignore":   liteTemplatesPath + ".prettierignore.tmpl",
		"eslint.config.mjs": liteTemplatesPath + "eslint.config.mjs.tmpl",

		typeutils.SrcPath + "index.ts":                  srcTmplPath + "index.ts.tmpl",
		routesPath + "index.ts":                         routesTmplPath + "index.ts.tmpl",
		controllersPath + "health/health.controller.ts": controllersTmplPath + "health/health.controller.ts.tmpl",
		configPath + "api.config.ts":                    configTmplPath + "api.config.ts.tmpl",
		configPath + "logger.config.ts":                 configTmplPath + "logger.config.ts.tmpl",
		middlewaresPath + "http-logger.middleware.ts":   middlewaresTmplPath + "http-logger.middleware.ts.tmpl",
		middlewaresPath + "error-handler.middleware.ts": middlewaresTmplPath + "error-handler.middleware.ts.tmpl",
		errorsPath + "ApiError.interface.ts":            errorsTmplPath + "ApiError.interface.ts.tmpl",
		errorsPath + "ApiError.ts":                      errorsTmplPath + "ApiError.ts.tmpl",
		errorsPath + "ApiErrorCode.enum.ts":             errorsTmplPath + "ApiErrorCode.enum.ts.tmpl",
		errorsPath + "ApiErrorKey.type.ts":              errorsTmplPath + "ApiErrorKey.type.ts.tmpl",
		errorsPath + "index.ts":                         errorsTmplPath + "index.ts.tmpl",
	}
}

func GetCleanFilesTemplatesPaths() map[string]string {
	cleanTmplPath := typeutils.GetCleanTemplatesPath + "/"

	adaptersPath, entitiesPath, frameworksPath, utilitiesPath :=
		typeutils.SrcPath+"adapters/", typeutils.SrcPath+"entities/", typeutils.SrcPath+"frameworks/", typeutils.SrcPath+"utilities/"

	adaptersTmplPath, entitiesTmplPath, frameworksTmplPath, utilitiesTmplPath :=
		cleanTmplPath+adaptersPath, cleanTmplPath+entitiesPath, cleanTmplPath+frameworksPath, cleanTmplPath+utilitiesPath

	testPath := "test/"

	return map[string]string{
		".gitignore":        cleanTmplPath + ".gitignore.tmpl",
		".prettierignore":   cleanTmplPath + ".prettierignore.tmpl",
		".prettierrc":       cleanTmplPath + ".prettierrc.tmpl",
		".env":              cleanTmplPath + ".env.tmpl",
		"eslint.config.mjs": cleanTmplPath + "eslint.config.mjs.tmpl",
		"package.json":      cleanTmplPath + "package.json.tmpl",
		"README.md":         cleanTmplPath + "README.md.tmpl",
		"tsconfig.json":     cleanTmplPath + "tsconfig.json.tmpl",
		"tsoa.json":         cleanTmplPath + "tsoa.json.tmpl",

		adaptersPath + "controllers/health/Health.controller.ts": adaptersTmplPath + "controllers/health/Health.controller.ts.tmpl",
		adaptersPath + "controllers/health/Health.dto.ts":        adaptersTmplPath + "controllers/health/Health.dto.ts.tmpl",
		adaptersPath + "services/Logger.service.ts":              adaptersTmplPath + "services/Logger.service.ts.tmpl",
		adaptersPath + "services/Health.service.ts":              adaptersTmplPath + "services/Health.service.ts.tmpl",
		adaptersPath + "middlewares/error-handler.middleware.ts": adaptersTmplPath + "middlewares/error-handler.middleware.ts.tmpl",
		adaptersPath + "middlewares/http-logger.middleware.ts":   adaptersTmplPath + "middlewares/http-logger.middleware.ts.tmpl",
		entitiesPath + "logger/Logger.interface.ts":              entitiesTmplPath + "logger/Logger.interface.ts.tmpl",
		entitiesPath + "error/ApiError.interface.ts":             entitiesTmplPath + "error/ApiError.interface.ts.tmpl",
		entitiesPath + "error/ApiError.ts":                       entitiesTmplPath + "error/ApiError.ts.tmpl",
		entitiesPath + "error/ApiErrorCode.enum.ts":              entitiesTmplPath + "error/ApiErrorCode.enum.ts.tmpl",
		entitiesPath + "error/ApiErrorKey.type.ts":               entitiesTmplPath + "error/ApiErrorKey.type.ts.tmpl",
		entitiesPath + "error/index.ts":                          entitiesTmplPath + "error/index.ts.tmpl",
		entitiesPath + "Health.ts":                               entitiesTmplPath + "Health.ts.tmpl",
		frameworksPath + "tsoa/services/iocContainer.ts":         frameworksTmplPath + "tsoa/services/iocContainer.ts.tmpl",
		frameworksPath + "tsoa/services/services.ts":             frameworksTmplPath + "tsoa/services/services.ts.tmpl",
		frameworksPath + "api.config.ts":                         frameworksTmplPath + "api.config.ts.tmpl",
		frameworksPath + "logger.config.ts":                      frameworksTmplPath + "logger.config.ts.tmpl",
		utilitiesPath + "di.constants.ts":                        utilitiesTmplPath + "di.constants.ts.tmpl",
		typeutils.SrcPath + "server.ts":                          cleanTmplPath + typeutils.SrcPath + "server.ts.tmpl",
		typeutils.SrcPath + "server_manager.ts":                  cleanTmplPath + typeutils.SrcPath + "server_manager.ts.tmpl",

		testPath + "test-server.ts": cleanTmplPath + testPath + "test-server.ts.tmpl",
	}
}

func GetTemplatePathsByProjectType(projectType typeutils.ProjectType) (map[string]string, error) {
	if projectType != typeutils.LiteProjectType && projectType != typeutils.CleanProjectType {
		return nil, fmt.Errorf("invalid project type: %s", projectType)
	}

	if projectType == typeutils.LiteProjectType {
		return GetLiteFilesTemplatesPaths(), nil
	}

	return GetCleanFilesTemplatesPaths(), nil
}

func CreateProjectStructureByType(projectPath string, projectType typeutils.ProjectType) error {
	logutils.Logger{}.Info(fmt.Sprintf("📂 Creating project structure for %s project...", projectType))
	if osutils.FileOrDirectoryExists(projectPath) {
		return fmt.Errorf(`❌ error generating structure, project already exists: %s`, projectPath)
	}

	projectStructure, err := GetProjectStructureByType(projectType)
	if err != nil {
		return err
	}

	for _, directory := range projectStructure {
		directoryPath := filepath.Join(projectPath, directory)
		if err := osutils.CreateDirectory(directoryPath); err != nil {
			return err
		}
	}

	return nil
}
