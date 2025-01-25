package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var NpmAPIUrl string = "https://registry.npmjs.org"

type PackageInfo struct {
	DistTags struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
}

var commonDependencies = DependencyMap{
	Dependencies: {"cors", "dotenv", "express", "uuid", "pino", "pino-http"},
	DevDependencies: {
		"@eslint/js", "@types/cors", "@types/express", "@types/node", "@types/uuid",
		"eslint", "pino-pretty", "prettier", "ts-node-dev", "tsconfig-paths",
		"typescript", "typescript-eslint",
	},
}

func cleanProjectDependencies() DependencyMap {
	dependencies := []string{"reflect-metadata", "swagger-ui-express", "tsoa", "tsyringe"}
	devDependencies := []string{
		"@types/jest", "@types/supertest", "@types/swagger-ui-express",
		"jest", "rimraf", "supertest", "ts-jest", "tsc-alias", "env-cmd",
	}

	return DependencyMap{
		Dependencies:    append(dependencies, commonDependencies[Dependencies]...),
		DevDependencies: append(devDependencies, commonDependencies[DevDependencies]...),
	}
}

func liteProjectDependencies() DependencyMap {
	dependencies := []string{"module-alias"}
	devDependencies := []string{}
	return DependencyMap{
		Dependencies:    append(dependencies, commonDependencies[Dependencies]...),
		DevDependencies: append(devDependencies, commonDependencies[DevDependencies]...),
	}
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

func getDependencyLts(dependency string, dependencyLst *string) error {
	dependencyUrl := fmt.Sprintf("%s/%s", NpmAPIUrl, dependency)
	res, err := http.Get(dependencyUrl)
	if err != nil {
		return err
	}
	defer func(response *http.Response) {
		if response.Body.Close(); err != nil {
			panic(err)
		}
	}(res)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var packageInfo PackageInfo
	if err := json.Unmarshal(body, &packageInfo); err != nil {
		return err
	}

	*dependencyLst = fmt.Sprintf("^%s", packageInfo.DistTags.Latest)

	return nil
}

func GetFormattedDependenciesByProjectType(projectType ProjectType) (FormattedDependencyMap, error) {
	if projectType != LiteProjectType && projectType != CleanProjectType {
		return nil, fmt.Errorf("invalid project type: %s", projectType)
	}

	formattedDependenciesMappedByProjectType := make(FormattedDependencyMap, 2)

	allDependencies := liteProjectDependencies()
	if projectType == CleanProjectType {
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
