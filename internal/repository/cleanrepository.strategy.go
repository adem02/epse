package repository

import "github.com/adem02/epse/internal/utils/typeutils"

type CleanRepositoryStrategy struct{}

func (s *CleanRepositoryStrategy) AddRepository(name string) (bool, error) {
	interfacePath := GetCleanRepositoryInterfaceFilePath(name)
	repoPath := GetCleanRepositoryFilePath(name)

	createdInterface, err := CreateRepositoryFileFromTmpl(typeutils.CleanProjectType, "repository.interface.ts.tmpl", name, interfacePath)
	if err != nil {
		return false, err
	}

	createdRepo, err := CreateRepositoryFileFromTmpl(typeutils.CleanProjectType, "repository.ts.tmpl", name, repoPath)
	if err != nil {
		return false, err
	}

	return createdInterface || createdRepo, nil
}
