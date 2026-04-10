package middleware

import (
	"fmt"

	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type MiddlewareManager struct {
	Name        string
	ProjectType typeutils.ProjectType
}

func NewMiddlewareManager(name string, projectType typeutils.ProjectType) *MiddlewareManager {
	return &MiddlewareManager{
		Name:        name,
		ProjectType: projectType,
	}
}

func (mm *MiddlewareManager) AddMiddleware() error {
	configData, err := config.ReadConfigFileData()
	if err != nil {
		return err
	}

	if MiddlewareAlreadyExists(configData.Middlewares, mm.Name) {
		return fmt.Errorf("middleware '%s' already exists", mm.Name)
	}

	if err := mm.createMiddlewareFile(); err != nil {
		return err
	}

	if err := config.AddNewMiddlewareInConfigFile(mm.Name, configData); err != nil {
		return err
	}

	mm.displaySuccess()

	return nil
}

func (mm *MiddlewareManager) createMiddlewareFile() error {
	templatePath := GetMiddlewareTemplateFilePath(mm.ProjectType)
	dirPath := GetMiddlewareDirectoryPath(mm.ProjectType)
	fileName := GetMiddlewareFileName(mm.Name)

	if err := osutils.CreateDirectory(dirPath); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", dirPath, fileName)
	tmplData := &struct{ MiddlewareName string }{
		MiddlewareName: mm.Name,
	}

	return osutils.CreateFileFromTmpl(templatePath, filePath, tmplData)
}

func (mm *MiddlewareManager) displaySuccess() {
	fmt.Println()
	logutils.Logger{}.Success("✅ Middleware added successfully!")
	fmt.Println()
	logutils.Logger{}.Info("📁 Files:")
	fmt.Printf("  ✓ %s/%s\n", GetMiddlewareDirectoryPath(mm.ProjectType), GetMiddlewareFileName(mm.Name))
	fmt.Printf("  ✓ epseconfig.json (updated)\n")
	fmt.Println()
	logutils.Logger{}.Info("📋 Usage:")

	if mm.ProjectType == typeutils.LiteProjectType {
		fmt.Printf("  import { %sMiddleware } from '@/middlewares/%s';\n\n", mm.Name, mm.Name+".middleware")
		fmt.Printf("  // On a specific route:\n")
		fmt.Printf("  router.get('/', %sMiddleware, YourController);\n\n", mm.Name)
		fmt.Printf("  // On all routes in a file:\n")
		fmt.Printf("  router.use(%sMiddleware);\n", mm.Name)
	} else {
		fmt.Printf("  import { %sMiddleware } from '@/adapters/middlewares/%s';\n\n", mm.Name, mm.Name+".middleware")
		fmt.Printf("  // On a specific endpoint:\n")
		fmt.Printf("  @Get('/')\n")
		fmt.Printf("  @Middlewares(%sMiddleware)\n", mm.Name)
		fmt.Printf("  async yourEndpoint(): Promise<void> { ... }\n")
	}

	fmt.Println()
}
