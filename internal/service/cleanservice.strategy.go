package service

import "github.com/adem02/epse/internal/utils/typeutils"

type CleanServiceStrategy struct{}

func (s *CleanServiceStrategy) AddService(names ServiceNames) (bool, error) {
	destPath := GetCleanServiceFilePath(names.FileName)
	return CreateServiceFileFromTmpl(typeutils.CleanProjectType, names.CleanName, destPath)
}
