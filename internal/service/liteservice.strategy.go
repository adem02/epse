package service

import (
	"strings"

	"github.com/adem02/epse/internal/utils/typeutils"
)

type LiteServiceStrategy struct{}

func (s *LiteServiceStrategy) AddService(names ServiceNames) (bool, error) {
	destPath := GetLiteServiceFilePath(strings.ToLower(names.FileName))
	return CreateServiceFileFromTmpl(typeutils.LiteProjectType, names.CleanName, destPath)
}
