package config

import (
	"fmt"
	"path/filepath"

	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

type ConfigManager struct {
}

type ConfigData struct {
	ProjectName     string                `json:"projectName"`
	ProjectType     string                `json:"projectType"`
	ControllersPath string                `json:"controllersPath"`
	Database        bool                  `json:"database"`
	AuthMiddleware  bool                  `json:"auth"`
	Routes          []typeutils.RouteData `json:"routes"`
}

var ConfigFilePath = filepath.Join(osutils.GetCurrentDirPath(), "epseconfig.json")

func GenerateNewConfigFile(projectType typeutils.ProjectType, projectName, projectPath string) error {
	logutils.Logger{}.Info("📝 Création du fichier de configuration...")
	if projectType != typeutils.LiteProjectType && projectType != typeutils.CleanProjectType {
		return fmt.Errorf("unable to create base config file. invalid project type: %s", projectType)
	}

	configData := ConfigData{
		ProjectName:     projectName,
		ProjectType:     string(projectType),
		ControllersPath: typeutils.ControllersPathMappedByProjectType[projectType],
		Routes: []typeutils.RouteData{
			typeutils.ConfigFileBaseRoute,
		},
	}

	configFilePath := filepath.Join(projectPath, "epseconfig.json")
	file, err := osutils.CreateFile(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return osutils.WriteJSONToFile(file, configData)
}

func ReadConfigFileData() (*ConfigData, error) {
	file, err := osutils.OpenFileWithWriteMode(ConfigFilePath, false)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	configData := &ConfigData{}
	if err := osutils.ParseJSONToStruct(file, configData); err != nil {
		return nil, err
	}

	return configData, nil
}

func AddNewRouteInConfigFile(domainName, routePrefix string, configData *ConfigData) error {
	configData.Routes = append(configData.Routes, typeutils.RouteData{
		Domaine:       domainName,
		RouteBasePath: "/" + routePrefix,
	})

	file, err := osutils.OpenFileWithWriteMode(ConfigFilePath, true)
	if err != nil {
		return err
	}
	defer file.Close()

	return osutils.WriteJSONToFile(file, configData)
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{}
}
