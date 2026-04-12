package repository

import "github.com/adem02/epse/internal/utils/typeutils"

type RepositoryStrategy interface {
	AddRepository(name string) (bool, error)
}

func GetRepositoryStrategy(projectType typeutils.ProjectType) RepositoryStrategy {
	switch projectType {
	case typeutils.LiteProjectType:
		return &LiteRepositoryStrategy{}
	case typeutils.CleanProjectType:
		return &CleanRepositoryStrategy{}
	default:
		panic("❌ Invalid project type")
	}
}
