package route

import "github.com/adem02/epse/internal/utils/typeutils"

type RouteStrategy interface {
	AddRoute(
		controllerNames ControllerNames,
		domainName string,
		completeRouteUrl string,
		method string,
		authMiddleware,
		validationMiddleware bool,
	) error
}

func GetRouteStrategy(projectType typeutils.ProjectType) RouteStrategy {
	switch projectType {
	case typeutils.LiteProjectType:
		return &LiteRouteStrategy{}
	case typeutils.CleanProjectType:
		return &CleanRouteStrategy{}
	default:
		panic("❌ Invalid project type")
	}
}
