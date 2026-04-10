package config

import (
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

func ConfigFileExists(path string) bool {
	return osutils.FileOrDirectoryExists(ConfigFilePath)
}

func DomainNameAlreadyExistsInRoutes(routes []typeutils.RouteData, domainName string) bool {
	for _, route := range routes {
		if route.Domaine == domainName {
			return true
		}
	}

	return false
}
