package repository

import (
	"fmt"
	"strings"

	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type RepositoryNames struct {
	CleanName                string
	LiteFileName             string
	LiteFileNameImportPath   string
	InterfaceFileName        string
	InterfaceImportPath      string
	ImplementationFileName   string
	ImplementationImportPath string
}

type RepositoryManager struct {
	Names       RepositoryNames
	ProjectType typeutils.ProjectType
}

func NewRepositoryManager(names RepositoryNames, projectType typeutils.ProjectType) *RepositoryManager {
	return &RepositoryManager{
		Names:       names,
		ProjectType: projectType,
	}
}

func (rm *RepositoryManager) AddRepository() error {
	strategy := GetRepositoryStrategy(rm.ProjectType)
	created, err := strategy.AddRepository(rm.Names)
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
		fmt.Printf("  ✓ src/repositories/%s\n", strings.ToLower(rm.Names.LiteFileName))
	} else {
		fmt.Printf("  ✓ src/useCases/gateway/%s\n", Capitalize(rm.Names.InterfaceFileName))
		fmt.Printf("  ✓ src/adapters/gateway/%s\n", Capitalize(rm.Names.ImplementationFileName))
	}

	fmt.Println()
}
