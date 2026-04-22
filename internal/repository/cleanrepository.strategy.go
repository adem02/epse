package repository

import "github.com/adem02/epse/internal/utils/typeutils"

type CleanRepositoryStrategy struct{}

func (s *CleanRepositoryStrategy) AddRepository(names RepositoryNames) (bool, error) {
	interfacePath := GetCleanRepositoryInterfaceFilePath(names.InterfaceFileName)
	repoPath := GetCleanRepositoryFilePath(names.ImplementationFileName)

	createdInterface, err := CreateRepositoryFileFromTmpl(
		typeutils.CleanProjectType,
		names,
		"repository.interface.ts.tmpl",
		interfacePath,
	)
	if err != nil {
		return false, err
	}

	createdRepo, err := CreateRepositoryFileFromTmpl(
		typeutils.CleanProjectType,
		names,
		"repository.ts.tmpl",
		repoPath,
	)
	if err != nil {
		return false, err
	}

	return createdInterface || createdRepo, nil
}
