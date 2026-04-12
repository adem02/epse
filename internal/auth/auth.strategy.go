package auth

import "github.com/adem02/epse/internal/utils/typeutils"

type AuthStrategy interface {
	AddAuth() error
}

func GetAuthStrategy(projectType typeutils.ProjectType) AuthStrategy {
	switch projectType {
	case typeutils.LiteProjectType:
		return &LiteAuthStrategy{}
	case typeutils.CleanProjectType:
		return &CleanAuthStrategy{}
	default:
		panic("❌ Invalid project type")
	}
}
