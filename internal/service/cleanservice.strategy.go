package service

import "github.com/adem02/epse/internal/utils/typeutils"

type CleanServiceStrategy struct{}

func (s *CleanServiceStrategy) AddService(name string) (bool, error) {
	destPath := GetCleanServiceFilePath(name)
	return CreateServiceFileFromTmpl(typeutils.CleanProjectType, name, destPath)
}
