package repository

import "github.com/adem02/epse/internal/utils/typeutils"

type LiteRepositoryStrategy struct{}

func (s *LiteRepositoryStrategy) AddRepository(name string) (bool, error) {
	destPath := GetLiteRepositoryFilePath(name)
	return CreateRepositoryFileFromTmpl(typeutils.LiteProjectType, "repository.ts.tmpl", name, destPath)
}
