package utils

import "fmt"

func GetLiteDirectoriesPaths(basePath string) []string {
	return []string{
		fmt.Sprintf("%s/src", basePath),
		fmt.Sprintf("%s/test", basePath),
		fmt.Sprintf("%s/src/controllers", basePath),
		fmt.Sprintf("%s/src/models", basePath),
		fmt.Sprintf("%s/src/middlewares", basePath),
		fmt.Sprintf("%s/src/repositories", basePath),
		fmt.Sprintf("%s/src/types", basePath),
		fmt.Sprintf("%s/src/services", basePath),
		fmt.Sprintf("%s/src/utils", basePath),
		fmt.Sprintf("%s/src/routes", basePath),
		fmt.Sprintf("%s/src/config", basePath),
	}
}

func GetCleanDirectoriesPaths(basePath string) []string {
	return []string{
		fmt.Sprintf("%s/test", basePath),
		fmt.Sprintf("%s/src", basePath),
		fmt.Sprintf("%s/src/adapters", basePath),
		fmt.Sprintf("%s/src/entities", basePath),
		fmt.Sprintf("%s/src/frameworks", basePath),
		fmt.Sprintf("%s/src/useCases", basePath),
		fmt.Sprintf("%s/src/utilities", basePath),
		fmt.Sprintf("%s/src/adapters/controllers", basePath),
		fmt.Sprintf("%s/src/adapters/controllers/health", basePath),
		fmt.Sprintf("%s/src/adapters/gateway", basePath),
		fmt.Sprintf("%s/src/adapters/middlewares", basePath),
		fmt.Sprintf("%s/src/adapters/services", basePath),
		fmt.Sprintf("%s/src/entities/error", basePath),
		fmt.Sprintf("%s/src/entities/types", basePath),
		fmt.Sprintf("%s/src/entities/logger", basePath),
		fmt.Sprintf("%s/src/frameworks/tsoa", basePath),
		fmt.Sprintf("%s/src/frameworks/tsoa/services", basePath),
	}
}
