package service

import (
	"fmt"
	"strings"

	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type ServiceManager struct {
	Name        string
	ProjectType typeutils.ProjectType
}

func NewServiceManager(name string, projectType typeutils.ProjectType) *ServiceManager {
	return &ServiceManager{
		Name:        name,
		ProjectType: projectType,
	}
}

func (sm *ServiceManager) AddService() error {
	sm.Name = CleanServiceName(sm.Name)

	strategy := GetServiceStrategy(sm.ProjectType)
	created, err := strategy.AddService(sm.Name)
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
		fmt.Printf("  ✓ src/services/%s.service.ts\n", strings.ToLower(sm.Name))
	} else {
		fmt.Printf("  ✓ src/adapters/services/%s.service.ts\n", Capitalize(sm.Name))
	}

	fmt.Println()
}
