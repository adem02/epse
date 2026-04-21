package middleware

import (
	"fmt"

	"github.com/adem02/epse/internal/config"
	"github.com/adem02/epse/internal/templates"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type MiddlewareNames struct {
	CleanName          string
	FileName           string
	FunctionName       string
	FileNameImportPath string
}

type MiddlewareManager struct {
	Names       MiddlewareNames
	ProjectType typeutils.ProjectType
}

func NewMiddlewareManager(names MiddlewareNames, projectType typeutils.ProjectType) *MiddlewareManager {
	return &MiddlewareManager{
		Names:       names,
		ProjectType: projectType,
	}
}

func (mm *MiddlewareManager) AddMiddleware() error {
	configData, err := config.ReadConfigFileData()
	if err != nil {
		return err
	}

	if MiddlewareAlreadyExists(configData.Middlewares, mm.Names.CleanName) {
		return fmt.Errorf("middleware '%s' already exists", mm.Names.CleanName)
	}

	if err := mm.createMiddlewareFile(); err != nil {
		return err
	}

	if err := config.AddNewMiddlewareInConfigFile(mm.Names.CleanName, configData); err != nil {
		return err
	}

	mm.displaySuccess()

	return nil
}

func (mm *MiddlewareManager) createMiddlewareFile() error {
	templatePath := GetMiddlewareTemplateFilePath(mm.ProjectType)
	dirPath := GetMiddlewareDirectoryPathByType(mm.ProjectType)
	fileName := mm.Names.FileName

	if err := osutils.CreateDirectory(dirPath); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", dirPath, fileName)
	tmplData := &struct{ MiddlewareName string }{
		MiddlewareName: mm.Names.CleanName,
	}

	return osutils.CreateFileFromTmpl(templates.FS, templatePath, filePath, tmplData)
}

func (mm *MiddlewareManager) displaySuccess() {
	fmt.Println()
	logutils.Logger{}.Success("✅ Middleware added successfully!")
	fmt.Println()
	logutils.Logger{}.Info("📁 Files:")
	fmt.Printf("  ✓ %s/%s\n", GetMiddlewareDirectoryPathByType(mm.ProjectType), mm.Names.FileName)
	fmt.Printf("  ✓ epseconfig.json (updated)\n")
	fmt.Println()
	logutils.Logger{}.Info("📋 Usage:")

	if mm.ProjectType == typeutils.LiteProjectType {
		fmt.Printf("  import { %s } from '@/middlewares/%s';\n\n", mm.Names.FunctionName, mm.Names.FileNameImportPath)
		fmt.Printf("  // On a specific route:\n")
		fmt.Printf("  router.get('/', %s, YourController);\n\n", mm.Names.FunctionName)
		fmt.Printf("  // On all routes in a file:\n")
		fmt.Printf("  router.use(%s);\n", mm.Names.FunctionName)
	} else {
		fmt.Printf("  import { %s } from '@/adapters/middlewares/%s';\n\n", mm.Names.FunctionName, mm.Names.FileNameImportPath)
		fmt.Printf("  // On a specific endpoint:\n")
		fmt.Printf("  @Get('/')\n")
		fmt.Printf("  @Middlewares(%s)\n", mm.Names.FunctionName)
		fmt.Printf("  async yourEndpoint(): Promise<void> { ... }\n")
	}

	fmt.Println()
}
