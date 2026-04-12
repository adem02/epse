package service

import "github.com/adem02/epse/internal/utils/typeutils"

type ServiceStrategy interface {
	AddService(name string) (bool, error)
}

func GetServiceStrategy(projectType typeutils.ProjectType) ServiceStrategy {
	switch projectType {
	case typeutils.LiteProjectType:
		return &LiteServiceStrategy{}
	case typeutils.CleanProjectType:
		return &CleanServiceStrategy{}
	default:
		panic("❌ Invalid project type")
	}
}
