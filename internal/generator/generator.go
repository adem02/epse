package generator

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/adem02/epse/internal/utils"
)

type Generator struct {
	ProjectName string
	ProjectType string
	Destination string
}

var title = color.New(color.Bold, color.FgCyan).SprintFunc()
var section = color.New(color.Bold, color.FgGreen).SprintFunc()
var info = color.New(color.FgHiBlue).SprintFunc()
var success = color.New(color.FgGreen).SprintFunc()
var warning = color.New(color.FgYellow).SprintFunc()

func generateLiteStructure(projectName, destination string) error {
	basePath := fmt.Sprintf("%s%s", destination, projectName)

	_, err := os.Stat(basePath)
	if !os.IsNotExist(err) {
		return errors.New("ℹ️ directory already exists")
	}

	for _, directory := range utils.GetLiteDirectoriesPaths(basePath) {
		err := os.MkdirAll(directory, os.ModePerm)

		if err != nil {
			if os.IsPermission(err) {
				fmt.Printf("❌ Permission insuffisante pour créer le répertoire : %s\n", directory)
			} else {
				fmt.Printf("❌ Erreur lors de la création du répertoire : %v\n", err)
			}
			return err
		}
	}

	templatesAbsPath, err := utils.GetAbsolutePath()

	if err != nil {
		return err
	}
	var outputFile *os.File

	for fileName, templatePath := range utils.GetLiteFilesTemplatesPaths() {
		tmplPath := filepath.Join(templatesAbsPath, templatePath)

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			return err
		}

		filePath := fmt.Sprintf("%s/%s", basePath, fileName)
		outputFile, err = os.Create(filePath)
		if err != nil {
			return err
		}

		err = tmpl.Execute(outputFile, struct{ ProjectName string }{ProjectName: projectName})
		if err != nil {
			return err
		}
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			panic(err)
		}
	}(outputFile)

	displayEndingMessage(projectName, destination, "lite")
	displayProjectDependenciesMessage(utils.GetLiteProjectDependencies())

	return nil
}

func generateCleanStructure(projectName, destination string) error {
	basePath := fmt.Sprintf("%s%s", destination, projectName)

	_, err := os.Stat(basePath)
	if !os.IsNotExist(err) {
		return errors.New("❌ directory already exists")
	}

	for _, directory := range utils.GetCleanDirectoriesPaths(basePath) {
		err := os.MkdirAll(directory, os.ModePerm)

		if err != nil {
			if os.IsPermission(err) {
				fmt.Printf("❌ Permission insuffisante pour créer le répertoire : %s\n", directory)
			} else {
				fmt.Printf("❌ Erreur lors de la création du répertoire : %v\n", err)
			}
			return err
		}
	}

	templatesAbsPath, err := utils.GetAbsolutePath()

	for fileName, templatePath := range utils.GetCleanFilesTemplatesPaths() {
		tmplPath := filepath.Join(templatesAbsPath, templatePath)

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			errorMessage := fmt.Errorf("error parsing template: %s", tmplPath)
			return errorMessage
		}

		filePath := fmt.Sprintf("%s/%s", basePath, fileName)

		outputFile, err := os.Create(filePath)
		if err != nil {
			errorMessage := fmt.Errorf("error creating file: %s", filePath)
			return errorMessage
		}

		defer func(outputFile *os.File) {
			err := outputFile.Close()
			if err != nil {
				panic(err)
			}
		}(outputFile)

		err = tmpl.Execute(outputFile, struct {
			ProjectName string
			EntryFile   string
			ApiPrefix   string
		}{
			ProjectName: projectName,
			EntryFile:   "./src/server.ts",
		})

		if err != nil {
			return err
		}
	}

	displayEndingMessage(projectName, destination, "clean")
	displayProjectDependenciesMessage(utils.GetCleanProjectDependencies())

	return nil
}

func (g Generator) GenerateProjectStructure() error {
	var err error = nil

	if g.ProjectType == "lite" {
		err = generateLiteStructure(g.ProjectName, g.Destination)
	} else if g.ProjectType == "clean" {
		err = generateCleanStructure(g.ProjectName, g.Destination)
	} else {
		return errors.New("failed to generate project structure of unknown type")
	}

	return err
}

func New(projectName, projectType, destination string) (Generator, error) {
	if projectType == "lite" {
		return Generator{
			ProjectName: projectName,
			ProjectType: projectType,
			Destination: destination,
		}, nil
	}

	if projectType == "clean" {
		return Generator{
			ProjectName: projectName,
			ProjectType: projectType,
			Destination: destination,
		}, nil
	}

	return Generator{}, errors.New("invalid project type")
}

func displayEndingMessage(projectName, destination, projectType string) {
	fmt.Println(success("✅ Génération réussie !"))
	fmt.Printf("   📂 %s: %s\n", section("Projet généré"), info(projectName))
	fmt.Printf("   📍 %s: %s\n", section("Emplacement"), info(destination))
	fmt.Printf("   🏗️ %s: %s\n\n", section("Type de projet"), info(projectType))
}

func displayProjectDependenciesMessage(allDependencies map[string][]string) {
	dependencies := strings.Join(allDependencies["dependencies"], " ")
	devDependencies := strings.Join(allDependencies["devDependencies"], " ")

	fmt.Println(warning("🚀 Étape suivante : Installez les dépendances"))
	fmt.Print("   Utilisez les commandes suivantes :\n\n")

	fmt.Println("   Avec npm :")
	fmt.Println(info(fmt.Sprintf("    npm install %s\n", dependencies)))
	fmt.Println(info(fmt.Sprintf("    npm install --save-dev %s\n", devDependencies)))

	fmt.Println("   Avec yarn :")
	fmt.Println(info(fmt.Sprintf("    yarn add %s\n", dependencies)))
	fmt.Println(info(fmt.Sprintf("    yarn add --dev %s\n", devDependencies)))

	fmt.Println(success("🎉 Votre projet est prêt ! Bon développement 🚀"))
}
