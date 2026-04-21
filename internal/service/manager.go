package service

import (
	"fmt"
	"strings"

	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type ServiceNames struct {
	CleanName          string
	FileName           string
	FileNameImportPath string
	FunctionName       string
}

type ServiceManager struct {
	Names       ServiceNames
	ProjectType typeutils.ProjectType
}

func NewServiceManager(names ServiceNames, projectType typeutils.ProjectType) *ServiceManager {
	return &ServiceManager{
		Names:       names,
		ProjectType: projectType,
	}
}

func (sm *ServiceManager) AddService() error {
	strategy := GetServiceStrategy(sm.ProjectType)
	created, err := strategy.AddService(sm.Names)
	if err != nil {
		return err
	}

	if created {
		sm.displaySuccess()
	}

	return nil
}

func (sm *ServiceManager) displaySuccess() {
	fmt.Println()
	logutils.Logger{}.Success("✅ Service added successfully!")
	fmt.Println()
	logutils.Logger{}.Info("📁 Files:")

	if sm.ProjectType == typeutils.LiteProjectType {
		fmt.Printf("  ✓ src/services/%s\n", strings.ToLower(sm.Names.FileName))
	} else {
		fmt.Printf("  ✓ src/adapters/services/%s\n", Capitalize(sm.Names.FileName))
	}

	fmt.Println()
}
