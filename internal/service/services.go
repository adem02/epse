package service

import (
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/common"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func GetServiceTemplatePath(projectType typeutils.ProjectType) string {
	return filepath.Join(
		"addcommand",
		string(projectType),
		"service",
		"service.ts.tmpl",
	)
}

func GetLiteServiceFilePath(name string) string {
	return common.GetFileOrDirectoryPathFromSrcPath("services", name)
}

func GetCleanServiceFilePath(name string) string {
	return common.GetFileOrDirectoryPathFromSrcPath("adapters", "services", name)
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func GenerateServiceNamesByType(name string, projectType typeutils.ProjectType) ServiceNames {
	words := normalizeServiceWords(name)

	pascalName := common.ToPascalCase(words)
	kebabName := common.ToKebabCase(words)

	fileBase := pascalName + ".service"
	if projectType == typeutils.LiteProjectType {
		fileBase = kebabName + ".service"
	}

	return ServiceNames{
		CleanName:          pascalName,
		FileName:           fileBase + ".ts",
		FileNameImportPath: fileBase,
		FunctionName:       pascalName + "Service",
	}
}

func normalizeServiceWords(input string) []string {
	s := strings.TrimSpace(input)
	if s == "" {
		return nil
	}

	lower := strings.ToLower(s)

	if strings.HasSuffix(lower, ".service.ts") {
		s = s[:len(s)-len(".service.ts")]
	} else if strings.HasSuffix(lower, ".ts") {
		s = s[:len(s)-len(".ts")]
	}

	lower = strings.ToLower(s)
	if strings.HasSuffix(lower, "-service") {
		s = s[:len(s)-len("-service")]
	} else if strings.HasSuffix(lower, "_service") {
		s = s[:len(s)-len("_service")]
	} else if strings.HasSuffix(lower, "service") {
		s = s[:len(s)-len("service")]
	}

	s = strings.TrimSpace(s)
	s = common.SplitCamelOrPascal(s)

	replacer := strings.NewReplacer("-", " ", "_", " ", ".", " ", "/", " ")
	s = replacer.Replace(s)

	return strings.Fields(strings.ToLower(s))
}
