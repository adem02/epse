package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func GetCurrentDirPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return cwd, nil
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

func CreateFileFromTmpl(tmplFilePath, fileName, basePath string, tmplData TmplData) error {
	currentDirectoryAbsolutePath, err := GetCurrentDirPath()
	if err != nil {
		return err
	}

	absoluteTmplFilePath := filepath.Join(currentDirectoryAbsolutePath, tmplFilePath)

	tmpl, err := template.ParseFiles(absoluteTmplFilePath)
	if err != nil {
		errorMessage := fmt.Errorf("❌ error parsing template: %s", absoluteTmplFilePath)
		return errorMessage
	}

	filePath := fmt.Sprintf("%s/%s", basePath, fileName)
	outputFile, err := os.Create(filePath)
	if err != nil {
		errorMessage := fmt.Errorf("❌ erreur lors de la création du fichier: %s", filePath)
		return errorMessage
	}
	defer func(outputFile *os.File) {
		if err := outputFile.Close(); err != nil {
			panic(err)
		}
	}(outputFile)

	if err := tmpl.Execute(outputFile, tmplData); err != nil {
		return err
	}

	return nil
}
