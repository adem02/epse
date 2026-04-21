package repository

import "github.com/adem02/epse/internal/utils/typeutils"

type LiteRepositoryStrategy struct{}

func (s *LiteRepositoryStrategy) AddRepository(names RepositoryNames) (bool, error) {
	destPath := GetLiteRepositoryFilePath(names.LiteFileName)
	return CreateRepositoryFileFromTmpl(
		typeutils.LiteProjectType,
		"repository.ts.tmpl",
		names.CleanName,
		destPath,
	)
}
