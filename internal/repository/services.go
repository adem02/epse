package repository

import (
	"path/filepath"
	"strings"

	"github.com/adem02/epse/internal/common"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func GetRepositoryTemplatePath(projectType typeutils.ProjectType, templateName string) string {
	return filepath.Join(
		"addcommand",
		string(projectType),
		"repository",
		templateName,
	)
}

func GetLiteRepositoryFilePath(fileName string) string {
	return common.GetFileOrDirectoryPathFromSrcPath("repositories", fileName)
}

func GetCleanRepositoryFilePath(repositoryFile string) string {
	return common.GetFileOrDirectoryPathFromSrcPath("adapters", "gateway", repositoryFile)
}

func GetCleanRepositoryInterfaceFilePath(repositoryInterfaceFile string) string {
	return common.GetFileOrDirectoryPathFromSrcPath("useCases", "gateway", repositoryInterfaceFile)
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func normalizeRepositoryWords(input string) []string {
	s := strings.TrimSpace(input)
	if s == "" {
		return nil
	}

	lower := strings.ToLower(s)

	if strings.HasSuffix(lower, ".repository.interface.ts") {
		s = s[:len(s)-len(".repository.interface.ts")]
	} else if strings.HasSuffix(lower, ".repository.ts") {
		s = s[:len(s)-len(".repository.ts")]
	} else if strings.HasSuffix(lower, ".ts") {
		s = s[:len(s)-len(".ts")]
	}

	lower = strings.ToLower(s)
	if strings.HasSuffix(lower, "-repository") {
		s = s[:len(s)-len("-repository")]
	} else if strings.HasSuffix(lower, "_repository") {
		s = s[:len(s)-len("_repository")]
	} else if strings.HasSuffix(lower, "repository") {
		s = s[:len(s)-len("repository")]
	}

	s = strings.TrimSpace(s)
	s = common.SplitCamelOrPascal(s)

	replacer := strings.NewReplacer("-", " ", "_", " ", ".", " ", "/", " ")
	s = replacer.Replace(s)

	return strings.Fields(strings.ToLower(s))
}

func GenerateRepositoryNamesByType(name string, projectType typeutils.ProjectType) RepositoryNames {
	words := normalizeRepositoryWords(name)

	pascalName := common.ToPascalCase(words)
	kebabName := common.ToKebabCase(words)

	return RepositoryNames{
		CleanName:                pascalName,
		LiteFileName:             kebabName + ".repository.ts",
		LiteFileNameImportPath:   kebabName + ".repository",
		InterfaceFileName:        pascalName + "Repository.interface.ts",
		InterfaceImportPath:      pascalName + "Repository.interface",
		ImplementationFileName:   pascalName + ".repository.ts",
		ImplementationImportPath: pascalName + ".repository",
	}
}
