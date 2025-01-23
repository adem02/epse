package utils

import "errors"

var directoriesPathsMappedByProjectType = map[ProjectType][]string{
	LiteProjectType: {
		"src",
		"test",
		"src/controllers",
		"src/models",
		"src/middlewares",
		"src/repositories",
		"src/types",
		"src/services",
		"src/utils",
		"src/routes",
		"src/config",
	},
	CleanProjectType: {
		"test",
		"src",
		"src/adapters",
		"src/entities",
		"src/frameworks",
		"src/useCases",
		"src/utilities",
		"src/adapters/controllers",
		"src/adapters/controllers/health",
		"src/adapters/gateway",
		"src/adapters/middlewares",
		"src/adapters/services",
		"src/entities/error",
		"src/entities/types",
		"src/entities/logger",
		"src/frameworks/tsoa",
		"src/frameworks/tsoa/services",
	},
}

func GetProjectStructureByType(projectType ProjectType) ([]string, error) {
	if projectType != LiteProjectType && projectType != CleanProjectType {
		return nil, errors.New("invalid project type")
	}

	return directoriesPathsMappedByProjectType[projectType], nil
}
