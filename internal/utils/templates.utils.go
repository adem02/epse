package utils

import (
	"os"
)

func GetAbsolutePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return cwd, nil
}

var GetLiteTemplatesPath = "templates/lite"
var GetCleanTemplatesPath = "templates/clean"
var SrcPath = "src/"

func GetLiteFilesTemplatesPaths() map[string]string {
	liteTemplatesPath := GetLiteTemplatesPath + "/"

	routesPath, configPath, middlewaresPath := SrcPath+"routes/", SrcPath+"config/", SrcPath+"middlewares/"
	srcTmplPath, routesTmplPath, configTmplPath, middlewaresTmplPath :=
		liteTemplatesPath+SrcPath, liteTemplatesPath+routesPath, liteTemplatesPath+configPath, liteTemplatesPath+middlewaresPath

	return map[string]string{
		"package.json":                  liteTemplatesPath + "package.json.tmpl",
		"README.md":                     liteTemplatesPath + "README.md.tmpl",
		"tsconfig.json":                 liteTemplatesPath + "tsconfig.json.tmpl",
		".env":                          liteTemplatesPath + ".env.tmpl",
		".prettierrc":                   liteTemplatesPath + ".prettierrc.tmpl",
		".prettierignore":               liteTemplatesPath + ".prettierignore.tmpl",
		"eslint.config.mjs":             liteTemplatesPath + "eslint.config.mjs.tmpl",
		SrcPath + "index.ts":            srcTmplPath + "index.ts.tmpl",
		routesPath + "index.ts":         routesTmplPath + "index.ts.tmpl",
		configPath + "api.config.ts":    configTmplPath + "api.config.ts.tmpl",
		configPath + "logger.config.ts": configTmplPath + "logger.config.ts.tmpl",
		middlewaresPath + "http_logger.middleware.ts": middlewaresTmplPath + "http_logger.middleware.ts.tmpl",
	}
}

func GetCleanFilesTemplatesPaths() map[string]string {
	cleanTmplPath := GetCleanTemplatesPath + "/"

	adaptersPath, entitiesPath, frameworksPath, utilitiesPath :=
		SrcPath+"adapters/", SrcPath+"entities/", SrcPath+"frameworks/", SrcPath+"utilities/"

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

		// Source files templates
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
		SrcPath + "server.ts":                                    cleanTmplPath + SrcPath + "server.ts.tmpl",
		SrcPath + "server_manager.ts":                            cleanTmplPath + SrcPath + "server_manager.ts.tmpl",

		// Test files templates
		testPath + "test-server.ts": cleanTmplPath + testPath + "test-server.ts.tmpl",
	}
}

func GetCleanProjectDependencies() map[string][]string {
	return map[string][]string{
		"dependencies": {
			"cors",
			"dotenv",
			"express",
			"pino",
			"pino-http",
			"reflect-metadata",
			"swagger-ui-express",
			"tsoa",
			"tsyringe",
			"uuid",
		},
		"devDependencies": {
			"@eslint/js",
			"@types/cors",
			"@types/express",
			"@types/jest",
			"@types/node",
			"@types/supertest",
			"@types/swagger-ui-express",
			"@types/uuid",
			"env-cmd",
			"eslint",
			"jest",
			"pino-pretty",
			"prettier",
			"rimraf",
			"supertest",
			"ts-jest",
			"ts-node-dev",
			"tsc-alias",
			"tsconfig-paths",
			"typescript",
			"typescript-eslint",
		},
	}
}

func GetLiteProjectDependencies() map[string][]string {
	return map[string][]string{
		"dependencies": {
			"cors",
			"dotenv",
			"express",
			"module-alias",
			"pino",
			"pino-http",
			"uuid",
		},
		"devDependencies": {
			"@eslint/js",
			"@types/cors",
			"@types/express",
			"@types/node",
			"eslint",
			"globals",
			"pino-pretty",
			"prettier",
			"ts-node-dev",
			"tsconfig-paths",
			"typescript",
			"typescript-eslint",
			"@types/uuid",
		},
	}
}
