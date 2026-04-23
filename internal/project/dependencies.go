package project

import (
	"fmt"
	"net/http"

	"github.com/adem02/epse/internal/utils/osutils"
	"github.com/adem02/epse/internal/utils/typeutils"
)

var NpmAPIUrl = "https://registry.npmjs.org"

type PackageInfo struct {
	DistTags struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
}

var commonDependencies = typeutils.DependencyMap{
	typeutils.Dependencies: {"cors", "dotenv", "express", "uuid", "pino", "pino-http"},
	typeutils.DevDependencies: {
		"@eslint/js", "@types/cors", "@types/express", "@types/node", "@types/uuid",
		"eslint", "pino-pretty", "prettier", "ts-node-dev", "tsconfig-paths",
		"typescript", "typescript-eslint",
	},
}

func cleanProjectDependencies() typeutils.DependencyMap {
	dependencies := []string{"reflect-metadata", "swagger-ui-express", "tsoa", "tsyringe"}
	devDependencies := []string{
		"@types/jest", "@types/supertest", "@types/swagger-ui-express",
		"jest", "rimraf", "supertest", "ts-jest", "tsc-alias", "env-cmd",
	}

	return typeutils.DependencyMap{
		typeutils.Dependencies:    append(dependencies, commonDependencies[typeutils.Dependencies]...),
		typeutils.DevDependencies: append(devDependencies, commonDependencies[typeutils.DevDependencies]...),
	}
}

func liteProjectDependencies() typeutils.DependencyMap {
	dependencies := []string{"module-alias"}
	devDependencies := []string{}

	return typeutils.DependencyMap{
		typeutils.Dependencies:    append(dependencies, commonDependencies[typeutils.Dependencies]...),
		typeutils.DevDependencies: append(devDependencies, commonDependencies[typeutils.DevDependencies]...),
	}
}

func getDependencyLts(dependency string, dependencyLst *string) error {
	dependencyUrl := fmt.Sprintf("%s/%s", NpmAPIUrl, dependency)
	res, err := http.Get(dependencyUrl)
	if err != nil {
		return err
		// return nil // In case of poor internet connectivity
	}
	defer res.Body.Close()

	packageInfo := &PackageInfo{}
	osutils.ParseJSONToStruct(res.Body, packageInfo)

	*dependencyLst = fmt.Sprintf("^%s", packageInfo.DistTags.Latest)

	return nil
}

func formatProjectDependencies(dependencies []string) (string, error) {
	var doneChans = make([]chan bool, len(dependencies))
	var errChans = make([]chan error, len(dependencies))
	dependencyResults := make([]string, len(dependencies))

	for i := 0; i < len(dependencies); i++ {
		doneChans[i] = make(chan bool)
		errChans[i] = make(chan error)
		go func(index int) {
			var dependencyLst string
			err := getDependencyLts(dependencies[index], &dependencyLst)
			if err != nil {
				errChans[index] <- err
				return
			}

			dependencyResults[index] = dependencyLst
			doneChans[index] <- true
		}(i)
	}

	for i := 0; i < len(dependencies); i++ {
		select {
		case err := <-errChans[i]:
			if err != nil {
				return "", err
			}
		case <-doneChans[i]:
		}
	}

	formattedDependencies := "{"
	for i, dependency := range dependencies {
		if i < len(dependencies)-1 {
			formattedDependencies += fmt.Sprintf(`
		"%s": "%s",`, dependency, dependencyResults[i])
		} else {
			formattedDependencies += fmt.Sprintf(`
		"%s": "%s"`, dependency, dependencyResults[i])
		}
	}
	formattedDependencies += `
	}`

	return formattedDependencies, nil
}

func GetFormattedDependenciesByProjectType(projectType typeutils.ProjectType) (typeutils.FormattedDependencyMap, error) {
	if projectType != typeutils.LiteProjectType && projectType != typeutils.CleanProjectType {
		return nil, fmt.Errorf("invalid project type: %s", projectType)
	}

	formattedDependenciesMappedByProjectType := make(typeutils.FormattedDependencyMap, 2)

	allDependencies := liteProjectDependencies()
	if projectType == typeutils.CleanProjectType {
		allDependencies = cleanProjectDependencies()
	}

	var err error
	for dependencyKey, dependenciesList := range allDependencies {
		formattedDependenciesMappedByProjectType[dependencyKey], err = formatProjectDependencies(dependenciesList)
		if err != nil {
			return nil, fmt.Errorf("failed to get formatted dependencies")
		}
	}

	return formattedDependenciesMappedByProjectType, nil
}
