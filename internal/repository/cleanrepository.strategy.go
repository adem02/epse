package repository

import "github.com/adem02/epse/internal/utils/typeutils"

type CleanRepositoryStrategy struct{}

func (s *CleanRepositoryStrategy) AddRepository(names RepositoryNames) (bool, error) {
	interfacePath := GetCleanRepositoryInterfaceFilePath(names.InterfaceFileName)
	repoPath := GetCleanRepositoryFilePath(names.ImplementationFileName)

	createdInterface, err := CreateRepositoryFileFromTmpl(
		typeutils.CleanProjectType,
		"repository.interface.ts.tmpl",
		names.CleanName,
		interfacePath,
	)
	if err != nil {
		return false, err
	}

	createdRepo, err := CreateRepositoryFileFromTmpl(
		typeutils.CleanProjectType,
		"repository.ts.tmpl",
		names.CleanName,
		repoPath,
	)
	if err != nil {
		return false, err
	}

	return createdInterface || createdRepo, nil
}
