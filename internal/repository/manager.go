package repository

import (
	"fmt"
	"strings"

	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type RepositoryManager struct {
	Name        string
	ProjectType typeutils.ProjectType
}

func NewRepositoryManager(name string, projectType typeutils.ProjectType) *RepositoryManager {
	return &RepositoryManager{
		Name:        name,
		ProjectType: projectType,
	}
}

func (rm *RepositoryManager) AddRepository() error {
	rm.Name = CleanRepositoryName(rm.Name)

	strategy := GetRepositoryStrategy(rm.ProjectType)
	created, err := strategy.AddRepository(rm.Name)
	if err != nil {
		return err
	}

	if created {
		rm.displaySuccess()
	}

	return nil
}

func (rm *RepositoryManager) displaySuccess() {
	fmt.Println()
	logutils.Logger{}.Success("✅ Repository added successfully!")
	fmt.Println()
	logutils.Logger{}.Info("📁 Files:")

	if rm.ProjectType == typeutils.LiteProjectType {
		fmt.Printf("  ✓ src/repositories/%s.repository.ts\n", strings.ToLower(rm.Name))
	} else {
		fmt.Printf("  ✓ src/useCases/gateway/%s.repository.interface.ts\n", Capitalize(rm.Name))
		fmt.Printf("  ✓ src/adapters/gateway/%s.repository.ts\n", Capitalize(rm.Name))
	}

	fmt.Println()
}
