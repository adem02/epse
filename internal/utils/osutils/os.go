package osutils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adem02/epse/internal/utils/logutils"
)

func GetCurrentDirPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		logutils.Logger{}.Error(fmt.Errorf("❌ impossible de récupérer le chemin du projet"))
		panic(err)
	}
	return cwd
}

func CreateDirectory(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		if os.IsPermission(err) {
			fmt.Printf("❌ insufficient permissions to create the directory: %s\n", path)
		} else {
			fmt.Printf("❌ error while creating the directory: %v\n", err)
		}
		return err
	}

	return nil
}

func GetCliRootPath() string {
	execPath, err := os.Executable()
	if err != nil {
		logutils.Logger{}.Error(fmt.Errorf("unable to get cli project path"))
		panic(err)
	}
	return filepath.Dir(execPath)
}

func FileOrDirectoryExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
