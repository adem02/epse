package service

import (
	"strings"

	"github.com/adem02/epse/internal/utils/typeutils"
)

type LiteServiceStrategy struct{}

func (s *LiteServiceStrategy) AddService(name string) (bool, error) {
	destPath := GetLiteServiceFilePath(strings.ToLower(name))
	return CreateServiceFileFromTmpl(typeutils.LiteProjectType, name, destPath)
}
