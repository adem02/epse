package middleware

import (
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/common"
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
	return filepath.Join("addcommand", string(projectType), "middleware", "custom.middleware.ts.tmpl")
}

func GetMiddlewareDirectoryPathByType(projectType typeutils.ProjectType) string {
	if projectType == typeutils.LiteProjectType {
		return common.GetFileOrDirectoryPathFromSrcPath("middlewares")
	}
	return common.GetFileOrDirectoryPathFromSrcPath("adapters", "middlewares")
}

func normalizeMiddlewareWords(input string) []string {
	s := strings.TrimSpace(input)
	if s == "" {
		return nil
	}

	lower := strings.ToLower(s)

	if strings.HasSuffix(lower, ".middleware.ts") {
		s = s[:len(s)-len(".middleware.ts")]
	} else if strings.HasSuffix(lower, ".ts") {
		s = s[:len(s)-len(".ts")]
	}

	lower = strings.ToLower(s)

	if strings.HasSuffix(lower, "-middleware") {
		s = s[:len(s)-len("-middleware")]
	} else if strings.HasSuffix(lower, "_middleware") {
		s = s[:len(s)-len("_middleware")]
	} else if strings.HasSuffix(lower, "middleware") {
		s = s[:len(s)-len("middleware")]
	}

	s = strings.TrimSpace(s)
	s = common.SplitCamelOrPascal(s)

	replacer := strings.NewReplacer("-", " ", "_", " ", ".", " ")
	s = replacer.Replace(s)

	return strings.Fields(strings.ToLower(s))
}

func GenerateMiddlewareNamesByType(name string, projectType typeutils.ProjectType) MiddlewareNames {
	words := normalizeMiddlewareWords(name)

	pascalName := common.ToPascalCase(words)
	kebabName := common.ToKebabCase(words)

	fileBase := pascalName + ".middleware"
	if projectType == typeutils.LiteProjectType {
		fileBase = kebabName + ".middleware"
	}

	return MiddlewareNames{
		CleanName:          pascalName,
		FileName:           fileBase + ".ts",
		FunctionName:       pascalName + "Middleware",
		FileNameImportPath: fileBase,
	}
}
