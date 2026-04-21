package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/adem02/epse/internal/common"
	"github.com/adem02/epse/internal/utils/logutils"
	"github.com/adem02/epse/internal/utils/osutils"
)

type CleanAuthStrategy struct{}

func (s *CleanAuthStrategy) AddAuth() error {
	if err := CreateAllCleanAuthFiles(); err != nil {
		return err
	}

	if err := UpdateApiErrorTypes(); err != nil {
		return err
	}

	if err := s.updateTsoaJson(); err != nil {
		return err
	}

	if err := UpdateDiConstants(); err != nil {
		return err
	}

	if err := osutils.AppendToFile(GetEnvFilePath(), "\n# JWT\nJWT_SECRET=your_secret_key_here\nJWT_EXPIRES_IN=24h\n"); err != nil {
		return err
	}

	return nil
}

func (s *CleanAuthStrategy) updateTsoaJson() error {
	tsoaJsonPath := common.GetFileOrDirectoryFromProjectRootPath("tsoa.json")
	bytes, err := os.ReadFile(tsoaJsonPath)
	if err != nil {
		return fmt.Errorf("❌ tsoa.json not found")
	}

	content := string(bytes)
	updated := false

	if !strings.Contains(content, "authenticationModule") {
		content = strings.Replace(
			content,
			`"routesDir"`,
			`"authenticationModule": "./src/frameworks/tsoa/authentication.ts",
    "routesDir"`,
			1,
		)
		updated = true
	}

	if !strings.Contains(content, "securityDefinitions") {
		content = strings.Replace(
			content,
			`"specVersion"`,
			`"securityDefinitions": {
      "jwt": {
        "type": "apiKey",
        "name": "Authorization",
        "in": "header"
      }
    },
    "specVersion"`,
			1,
		)
		updated = true
	}

	if !updated {
		logutils.Logger{}.Warning("⚠️ tsoa.json already up to date")
		return nil
	}

	return os.WriteFile(tsoaJsonPath, []byte(content), 0644)
}
